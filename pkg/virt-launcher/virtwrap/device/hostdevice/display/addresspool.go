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
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/log"

	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/device/hostdevice"
)

// NewDisplayAddressPool creates a Display address pool based on the provided list of host-devices and
// the environment variables that describe the resource.
func NewDisplayAddressPool(displayDevices []v1.Display) *hostdevice.AddressPool {
	return hostdevice.NewAddressPool(v1.DisplayResourcePrefix, extractResources(displayDevices))
}

func extractResources(displayDevices []v1.Display) []string {
	log.Log.Warningf("===== extractResources")
	var resourceSet = make(map[string]struct{})
	for _, device := range displayDevices {
		log.Log.Warningf("===== displayDevice.DeviceName %s", device.DeviceName)
		resourceSet[device.DeviceName] = struct{}{}
		log.Log.Warningf("===== resourceSet =%v", resourceSet)
	}

	var resources []string
	for resource, _ := range resourceSet {
		resources = append(resources, resource)
	}
	log.Log.Warningf("===== returning resources =%v", resources)
	return resources
}
