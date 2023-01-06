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
 * Copyright 2021 Red Hat, Inc.
 *
 */

package display

import (
	"fmt"

	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/log"

	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
)

type DisplayDeviceMetaData struct {
	AliasPrefix  string
	Name         string
	ResourceName string
	Monitor      string
}

//type createDisplayDevice func(DisplayDeviceMetaData, string) (*[]api.Arg, *[]api.Env, error)

type AddressPooler interface {
	Pop(key string) (value string, err error)
}

// func CreateDisplayDevices(devicesData []DisplayDeviceMetaData, DisplayAddrPool AddressPooler) ([]api.Arg, []api.Env, error) {
// 	return createDisplayDevices(devicesData, DisplayAddrPool)
// }

const (
	failedCreateDisplayDeviceFmt = "failed to create Display devices: %v"
	AliasPrefix                  = "display-"
	DefaultDisplayOn             = true
)

func CreateDisplayDevices(vmiDisplays []v1.Display) (api.Commandline, error) {
	return CreateDisplayDevicesFromPools(vmiDisplays, NewDisplayAddressPool(vmiDisplays))
}

func CreateDisplayDevicesFromPools(vmiDisplays []v1.Display, displayAddressPool AddressPooler) (api.Commandline, error) {
	pool := NewBestEffortAddressPool(displayAddressPool)

	devicesMetaData := createDisplayDevicesMetadata(vmiDisplays)
	displayCmd, err := createDisplayDevices(devicesMetaData, pool)
	if err != nil {
		return api.Commandline{}, fmt.Errorf(failedCreateDisplayDeviceFmt, err)
	}

	// if err := validateCreationOfAllDevices(vmiDisplays, displayarg); err != nil {
	// 	log.Log.Warningf("===== hacking to return devices")
	// 	return nil, nil, fmt.Errorf(failedCreateDisplayDeviceFmt, err)
	// }

	return displayCmd, nil
}

func createDisplayDevicesMetadata(vmiDisplays []v1.Display) []DisplayDeviceMetaData {
	var devicesMetaData []DisplayDeviceMetaData
	for _, dev := range vmiDisplays {
		devicesMetaData = append(devicesMetaData, DisplayDeviceMetaData{
			AliasPrefix:  AliasPrefix,
			Name:         dev.Name,
			ResourceName: dev.DeviceName,
			Monitor:      dev.Monitor,
		})
	}
	return devicesMetaData
}

// validateCreationOfAllDevices validates that all specified Display/s have a matching host-device.
// On validation failure, an error is returned.
// The validation assumes that the assignment of a device to a specified Display is correct,
// therefore a simple quantity check is sufficient.
func validateCreationOfAllDevices(displays []v1.Display, hostDevices []api.Arg) error {
	if len(displays) != len(hostDevices) {
		return fmt.Errorf(
			"the number of display/s do not match the number of devices:\nDisplay: %v\nDevice: %v", displays, hostDevices,
		)
	}
	return nil
}

func createDisplayDevices(devicesData []DisplayDeviceMetaData, addrPool AddressPooler) (api.Commandline, error) {
	var displayArg []api.Arg
	var displayEnv []api.Env

	log.Log.Infof("==========display device data: %v", devicesData)

	for _, deviceData := range devicesData {
		address, err := addrPool.Pop(deviceData.ResourceName)
		if err != nil {
			return api.Commandline{}, fmt.Errorf(failedCreateDisplayDeviceFmt, err)
		}
		log.Log.Infof("==========   address: %s", address)

		// if pop succeeded but the address is empty, ignore the device and let the caller
		// decide if this is acceptable or not.
		if address == "" {
			log.Log.Infof("==========empty display address: %v", devicesData)
			continue
		}

		displayarg, displayenv, err := createDisplayDev(deviceData, address, deviceData.Monitor)
		if err != nil {
			return api.Commandline{}, fmt.Errorf(failedCreateDisplayDeviceFmt, err)
		}

		displayArg = append(displayArg, api.Arg{Value: "-display"}, *displayarg)
		displayEnv = append(displayEnv, *displayenv)
		log.Log.Infof("host-device created: %s", address)
	}

	displayCmd := api.Commandline{
		QEMUArg: displayArg,
		QEMUEnv: displayEnv,
	}

	return displayCmd, nil
}

// flagXY
func createDisplayDev(displayDev DisplayDeviceMetaData, displayDevAddr string, monitor string) (*api.Arg, *api.Env, error) {
	// displayDevAddr = $DISPLAY.MONITOR eg : :1.0
	// split := strings.Split(displayDevAddr, ".")
	// busdevNum = addrPool.Pop(hostDeviceData.ResourceName) = [hostDeviceData.ResourceName].value
	log.Log.Infof("==========   generate displayArg address: %s", displayDevAddr)
	var displayArg api.Arg

	displayArg = api.Arg{Value: fmt.Sprintf("gtk,gl=on,full-screen=on,monitor.0=%s", monitor)}

	var displayEnv api.Env
	displayEnv = api.Env{Name: "DISPLAY", Value: ":1"}

	return &displayArg, &displayEnv, nil
}
