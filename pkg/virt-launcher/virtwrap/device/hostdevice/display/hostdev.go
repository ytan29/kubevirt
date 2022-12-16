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
	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/device/hostdevice"
)

const (
	failedCreateDisplayHostDeviceFmt = "failed to create Display host-devices: %v"
	AliasPrefix                      = "display-"
	DefaultDisplayOn                 = true
)

func CreateHostDevices(vmiDisplays []v1.Display) ([]api.HostDevice, error) {
	return CreateHostDevicesFromPools(vmiDisplays, NewDisplayAddressPool(vmiDisplays))
}

func CreateHostDevicesFromPools(vmiDisplays []v1.Display, displayAddressPool hostdevice.AddressPooler) ([]api.HostDevice, error) {
	pool := hostdevice.NewBestEffortAddressPool(displayAddressPool)

	hostDevicesMetaData := createHostDevicesMetadata(vmiDisplays)
	hostDevices, err := hostdevice.CreateDisplayHostDevices(hostDevicesMetaData, pool)
	if err != nil {
		return nil, fmt.Errorf(failedCreateDisplayHostDeviceFmt, err)
	}

	if err := validateCreationOfAllDevices(vmiDisplays, hostDevices); err != nil {
		log.Log.Warningf("===== hacking to return devices")
		return nil, fmt.Errorf(failedCreateDisplayHostDeviceFmt, err)
	}

	return hostDevices, nil
}

func createHostDevicesMetadata(vmiDisplays []v1.Display) []hostdevice.HostDeviceMetaData {
	var hostDevicesMetaData []hostdevice.HostDeviceMetaData
	for _, dev := range vmiDisplays {
		hostDevicesMetaData = append(hostDevicesMetaData, hostdevice.HostDeviceMetaData{
			AliasPrefix:  AliasPrefix,
			Name:         dev.Name,
			ResourceName: dev.DeviceName,
		})
	}
	// sample
	// AliasPrefix:  display-
	// Name:         displaydev
	// ResourceName: generic/monitor
	return hostDevicesMetaData
}

// validateCreationOfAllDevices validates that all specified Display/s have a matching host-device.
// On validation failure, an error is returned.
// The validation assumes that the assignment of a device to a specified Display is correct,
// therefore a simple quantity check is sufficient.
func validateCreationOfAllDevices(displays []v1.Display, hostDevices []api.HostDevice) error {
	if len(displays) != len(hostDevices) {
		return fmt.Errorf(
			"the number of display/s do not match the number of devices:\nDisplay: %v\nDevice: %v", displays, hostDevices,
		)
	}
	return nil
}
