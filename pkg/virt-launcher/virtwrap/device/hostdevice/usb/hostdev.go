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

package usb

import (
	"fmt"

	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/log"

	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/device/hostdevice"
)

const (
	failedCreateUSBHostDeviceFmt = "failed to create USB host-devices: %v"
	AliasPrefix                  = "usb-"
	DefaultDisplayOn             = true
)

func CreateHostDevices(vmiUSBs []v1.USB) ([]api.HostDevice, error) {
	return CreateHostDevicesFromPools(vmiUSBs, NewUSBAddressPool(vmiUSBs))
}

func CreateHostDevicesFromPools(vmiUSBs []v1.USB, usbAddressPool hostdevice.AddressPooler) ([]api.HostDevice, error) {
	usbPool := hostdevice.NewBestEffortAddressPool(usbAddressPool)

	hostDevicesMetaData := createHostDevicesMetadata(vmiUSBs)
	usbHostDevices, err := hostdevice.CreateUSBHostDevices(hostDevicesMetaData, usbPool)
	if err != nil {
		return nil, fmt.Errorf(failedCreateUSBHostDeviceFmt, err)
	}

	if err := validateCreationOfAllDevices(vmiUSBs, usbHostDevices); err != nil {
		log.Log.Warningf("===== hacking to return devices")
		return nil, fmt.Errorf(failedCreateUSBHostDeviceFmt, err)
	}

	return usbHostDevices, nil
}

func createHostDevicesMetadata(vmiUSBs []v1.USB) []hostdevice.HostDeviceMetaData {
	var hostDevicesMetaData []hostdevice.HostDeviceMetaData
	for _, dev := range vmiUSBs {
		hostDevicesMetaData = append(hostDevicesMetaData, hostdevice.HostDeviceMetaData{
			AliasPrefix:  AliasPrefix,
			Name:         dev.Name,
			ResourceName: dev.DeviceName,
			BusPort:      dev.BusPort,
		})
	}
	// sample
	// AliasPrefix:  usb-
	// Name:         usbdev
	// ResourceName: generic/hid-mouse
	// BusPort: "1-9"
	return hostDevicesMetaData
}

// validateCreationOfAllDevices validates that all specified USB/s have a matching host-device.
// On validation failure, an error is returned.
// The validation assumes that the assignment of a device to a specified USB is correct,
// therefore a simple quantity check is sufficient.
func validateCreationOfAllDevices(usbs []v1.USB, hostDevices []api.HostDevice) error {
	if len(usbs) != len(hostDevices) {
		return fmt.Errorf(
			"the number of USB/s do not match the number of devices:\nUSB: %v\nDevice: %v", usbs, hostDevices,
		)
	}
	return nil
}
