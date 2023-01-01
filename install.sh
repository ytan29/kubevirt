#!/bin/bash
kubectl get pods -n kubevirt | grep virt-operator
if [ $? -eq 1 ] ; then
	kubectl create -f _out/manifests/release/kubevirt-operator.yaml && kubectl create -f _out/manifests/release/kubevirt-cr.yaml
else
	echo "found kubevirt in cluster. failed to install"
fi
