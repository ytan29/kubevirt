/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2018 Red Hat, Inc.
 *
 */

package device_manager

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"kubevirt.io/client-go/log"

	v1 "kubevirt.io/api/core/v1"

	"kubevirt.io/kubevirt/pkg/util"
	pluginapi "kubevirt.io/kubevirt/pkg/virt-handler/device-manager/deviceplugin/v1beta1"
)

const (
	usbDevicePath = "/dev/bus/usb"
	usbBasePath   = "/sys/bus/usb/devices"
)

type USBDevice struct {
	VendorID    string
	ProductID   string
	BusNum      string
	DevNum      string
	FullDevPath string
}

type USBDevicePlugin struct {
	devs         []*pluginapi.Device
	server       *grpc.Server
	socketPath   string
	stop         <-chan struct{}
	health       chan deviceHealth
	devicePath   string
	resourceName string
	done         chan struct{}
	deviceRoot   string
	initialized  bool
	lock         *sync.Mutex
	deregistered chan struct{}
}

func NewUSBDevicePlugin(usbDevices []*USBDevice, resourceName string) *USBDevicePlugin {
	serverSock := SocketPath(strings.Replace(resourceName, "/", "-", -1))
	log.DefaultLogger().Infof("=========== NewUSBDevicePlugin socketPath %s", serverSock)

	initHandler()

	devs := constructDPIdevicesFromUsb(usbDevices)
	dpi := &USBDevicePlugin{
		devs:         devs,
		socketPath:   serverSock,
		resourceName: resourceName,
		devicePath:   usbDevicePath,
		deviceRoot:   util.HostRootMount,
		health:       make(chan deviceHealth),
		initialized:  false,
		lock:         &sync.Mutex{},
	}
	return dpi
}

func constructDPIdevicesFromUsb(usbDevices []*USBDevice) (devs []*pluginapi.Device) {
	for _, usbDevice := range usbDevices {
		dpiDev := &pluginapi.Device{
			ID:     usbDevice.BusNum + "/" + usbDevice.DevNum,
			Health: pluginapi.Healthy,
		}
		devs = append(devs, dpiDev)
	}
	return
}

// Start starts the device plugin
func (dpi *USBDevicePlugin) Start(stop <-chan struct{}) (err error) {
	log.Log.Info("===================start usb===================")
	logger := log.DefaultLogger()
	dpi.stop = stop
	dpi.done = make(chan struct{})
	dpi.deregistered = make(chan struct{})

	err = dpi.cleanup()
	if err != nil {
		return err
	}

	sock, err := net.Listen("unix", dpi.socketPath)
	if err != nil {
		return fmt.Errorf("error creating GRPC server socket: %v", err)
	}

	dpi.server = grpc.NewServer([]grpc.ServerOption{}...)
	defer dpi.stopDevicePlugin()

	pluginapi.RegisterDevicePluginServer(dpi.server, dpi)

	errChan := make(chan error, 2)

	go func() {
		errChan <- dpi.server.Serve(sock)
	}()

	err = waitForGRPCServer(dpi.socketPath, connectionTimeout)
	if err != nil {
		return fmt.Errorf("error starting the GRPC server: %v", err)
	}

	err = dpi.register()
	if err != nil {
		return fmt.Errorf("error registering with device plugin manager: %v", err)
	}

	go func() {
		errChan <- dpi.healthCheck()
	}()

	dpi.setInitialized(true)
	logger.Infof("%s device plugin started", dpi.resourceName)
	err = <-errChan

	return err
}

func (dpi *USBDevicePlugin) ListAndWatch(_ *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
	log.Log.Info("===================list and watch usb===================")
	s.Send(&pluginapi.ListAndWatchResponse{Devices: dpi.devs})

	done := false
	for {
		select {
		case devHealth := <-dpi.health:
			for _, dev := range dpi.devs {
				if devHealth.DevId == dev.ID {
					dev.Health = devHealth.Health
				}
			}
			log.Log.Info("health status s.Send(&pluginapi.ListAndWatchResponse{Devices: dpi.devs})")
			s.Send(&pluginapi.ListAndWatchResponse{Devices: dpi.devs})
		case <-dpi.stop:
			done = true
		case <-dpi.done:
			done = true
		}
		if done {
			break
		}
	}
	// Send empty list to increase the chance that the kubelet acts fast on stopped device plugins
	// There exists no explicit way to deregister devices
	emptyList := []*pluginapi.Device{}
	if err := s.Send(&pluginapi.ListAndWatchResponse{Devices: emptyList}); err != nil {
		log.DefaultLogger().Reason(err).Infof("%s device plugin failed to deregister", dpi.resourceName)
	}
	close(dpi.deregistered)
	return nil
}

func (dpi *USBDevicePlugin) Allocate(_ context.Context, r *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	log.Log.Info("===================allocate usb===================")
	log.DefaultLogger().Infof("Allocate: resourceName: %s", dpi.resourceName)
	resourceNameEnvVar := util.ResourceNameToEnvVar(v1.USBResourcePrefix, dpi.resourceName)
	log.DefaultLogger().Infof("Allocate: resourceNameEnvVar: %s", resourceNameEnvVar)
	allocatedDevices := []string{}
	resp := new(pluginapi.AllocateResponse)
	containerResponse := new(pluginapi.ContainerAllocateResponse)

	for _, request := range r.ContainerRequests {
		log.DefaultLogger().Infof("Allocate: request: %v", request)
		deviceSpecs := make([]*pluginapi.DeviceSpec, 0)
		for _, devID := range request.DevicesIDs {
			log.DefaultLogger().Infof("Allocate: devID: %s", devID)
			allocatedDevices = append(allocatedDevices, dpi.resourceName)

			formattedUSB := formatUSBDeviceSpecs(devID)
			log.DefaultLogger().Infof("Allocate: formatted usb: %v", formattedUSB)

			deviceSpecs = append(deviceSpecs, formattedUSB...)
		}
		containerResponse.Devices = deviceSpecs
		envVar := make(map[string]string)
		envVar[resourceNameEnvVar] = strings.Join(allocatedDevices, ",")

		containerResponse.Envs = envVar
		resp.ContainerResponses = append(resp.ContainerResponses, containerResponse)
	}
	return resp, nil
}

func (dpi *USBDevicePlugin) healthCheck() error {
	log.Log.Info("===================healthcheck for usb===================")

	logger := log.DefaultLogger()
	monitoredDevices := make(map[string]string)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to creating a fsnotify watcher: %v", err)
	}
	defer watcher.Close()

	// This way we don't have to mount /dev from the node
	devicePath := filepath.Join(dpi.deviceRoot, dpi.devicePath)

	// Start watching the files before we check for their existence to avoid races
	dirName := filepath.Dir(devicePath)
	err = watcher.Add(dirName)
	if err != nil {
		return fmt.Errorf("failed to add the device root path to the watcher: %v", err)
	}

	_, err = os.Stat(devicePath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("could not stat the device: %v", err)
		}
		logger.Warningf("device '%s' is not present, the device plugin can't expose it.", dpi.devicePath)
		dpi.health <- deviceHealth{Health: pluginapi.Unhealthy}
	}

	// probe all devices
	for _, dev := range dpi.devs {
		usbDevice := filepath.Join(devicePath, dev.ID)
		err = watcher.Add(usbDevice)
		logger.Infof("watching %s", usbDevice)
		if err != nil {
			return fmt.Errorf("failed to add the device %s to the watcher: %v", usbDevice, err)
		}
		monitoredDevices[usbDevice] = dev.ID
	}
	logger.Infof("device '%s' is present.", dpi.devicePath)

	dirName = filepath.Dir(dpi.socketPath)
	err = watcher.Add(dirName)

	if err != nil {
		return fmt.Errorf("failed to add the device-plugin kubelet path to the watcher: %v", err)
	}
	_, err = os.Stat(dpi.socketPath)
	if err != nil {
		return fmt.Errorf("failed to stat the device-plugin socket: %v", err)
	}

	for {
		select {
		case <-dpi.stop:
			return nil
		case err := <-watcher.Errors:
			logger.Reason(err).Errorf("error watching devices and device plugin directory")
		case event := <-watcher.Events:
			logger.V(4).Infof("health Event: %v", event)
			// if monDevId, exist := monitoredDevices[event.Name]; exist {
			// 	// Health in this case is if the device path actually exists
			// 	if event.Op == fsnotify.Create {
			// 		logger.Infof("monitored device %s appeared", dpi.resourceName)
			// 		dpi.health <- deviceHealth{
			// 			DevId:  monDevId,
			// 			Health: pluginapi.Healthy,
			// 		}
			// 	} else if (event.Op == fsnotify.Remove) || (event.Op == fsnotify.Rename) {
			// 		logger.Infof("monitored device %s disappeared", dpi.resourceName)
			// 		dpi.health <- deviceHealth{
			// 			DevId:  monDevId,
			// 			Health: pluginapi.Unhealthy,
			// 		}
			// 	}
			if event.Name == devicePath {
				// Health in this case is if the device path actually exists
				if event.Op == fsnotify.Create {
					logger.Infof("monitored device %s appeared", dpi.resourceName)
					dpi.health <- deviceHealth{Health: pluginapi.Healthy}
				} else if (event.Op == fsnotify.Remove) || (event.Op == fsnotify.Rename) {
					logger.Infof("monitored device %s disappeared", dpi.resourceName)
					dpi.health <- deviceHealth{Health: pluginapi.Unhealthy}
				}
			} else if event.Name == dpi.socketPath && event.Op == fsnotify.Remove {
				logger.Infof("device socket file for device %s was removed, kubelet probably restarted.", dpi.resourceName)
				return nil
			} else {
				logger.Infof("something else....")
			}
		}
	}
}

func (dpi *USBDevicePlugin) GetDevicePath() string {
	return dpi.devicePath
}

func (dpi *USBDevicePlugin) GetDeviceName() string {
	return dpi.resourceName
}

// Stop stops the gRPC server
func (dpi *USBDevicePlugin) stopDevicePlugin() error {
	defer func() {
		if !IsChanClosed(dpi.done) {
			close(dpi.done)
		}
	}()

	// Give the device plugin one second to properly deregister
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	select {
	case <-dpi.deregistered:
	case <-ticker.C:
	}

	dpi.server.Stop()
	dpi.setInitialized(false)
	return dpi.cleanup()
}

// Register registers the device plugin for the given resourceName with Kubelet.
func (dpi *USBDevicePlugin) register() error {
	conn, err := gRPCConnect(pluginapi.KubeletSocket, connectionTimeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pluginapi.NewRegistrationClient(conn)
	reqt := &pluginapi.RegisterRequest{
		Version:      pluginapi.Version,
		Endpoint:     path.Base(dpi.socketPath),
		ResourceName: dpi.resourceName,
	}

	_, err = client.Register(context.Background(), reqt)
	if err != nil {
		return err
	}
	return nil
}

func (dpi *USBDevicePlugin) cleanup() error {
	if err := os.Remove(dpi.socketPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (dpi *USBDevicePlugin) GetDevicePluginOptions(_ context.Context, _ *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	options := &pluginapi.DevicePluginOptions{
		PreStartRequired: false,
	}
	return options, nil
}

func (dpi *USBDevicePlugin) PreStartContainer(_ context.Context, _ *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	res := &pluginapi.PreStartContainerResponse{}
	return res, nil
}

func discoverPermittedHostUSBDevices(supportedUSBDeviceMap map[string]string) map[string][]*USBDevice {
	initHandler()

	usbDevicesMap := make(map[string][]*USBDevice)

	for addr, resourceName := range supportedUSBDeviceMap {
		addarr := strings.Split(addr, "-")

		usbdev := &USBDevice{
			BusNum: addarr[0],
			DevNum: addarr[1],
		}

		// Confirm the path is valid and thus device is available
		fullpathToCheck := filepath.Join(usbDevicePath, usbdev.BusNum, usbdev.DevNum)
		_, err := os.Stat(fullpathToCheck)
		if err != nil {
			log.DefaultLogger().Reason(err).Errorf("failed identify any usb device at %s", fullpathToCheck)
			return nil
		}

		usbdev.FullDevPath = fullpathToCheck
		log.DefaultLogger().V(4).Infof("Found usb device at %s", usbdev.FullDevPath)

		usbDevicesMap[resourceName] = append(usbDevicesMap[resourceName], usbdev)
	}
	return usbDevicesMap
}

func (dpi *USBDevicePlugin) GetInitialized() bool {
	dpi.lock.Lock()
	defer dpi.lock.Unlock()
	return dpi.initialized
}

func (dpi *USBDevicePlugin) setInitialized(initialized bool) {
	dpi.lock.Lock()
	dpi.initialized = initialized
	dpi.lock.Unlock()
}
