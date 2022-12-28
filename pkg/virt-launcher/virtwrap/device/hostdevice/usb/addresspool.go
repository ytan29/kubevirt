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
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/log"

	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/device/hostdevice"
)

// NewUSBAddressPool creates a USB address pool based on the provided list of host-devices and
// the environment variables that describe the resource.
func NewUSBAddressPool(usbDevices []v1.USB) *hostdevice.AddressPool {
	return hostdevice.NewAddressPool(v1.USBResourcePrefix, extractResources(usbDevices))
}

func extractResources(usbDevices []v1.USB) []string {
	var resourceSet = make(map[string]struct{})
	for _, usbDevice := range usbDevices {
		log.Log.Warningf("===== usbDevice.DeviceName %s", usbDevice.DeviceName)
		resourceSet[usbDevice.DeviceName] = struct{}{}
		log.Log.Warningf("===== resourceSet =%v", resourceSet)
	}

	var resources []string
	for resource, _ := range resourceSet {
		resources = append(resources, resource)
	}
	log.Log.Warningf("===== returning resources =%v", resources)
	return resources
}
