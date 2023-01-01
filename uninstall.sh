#!/bin/bash
kubectl get pods -n kubevirt | grep virt-operator
if [ $? -eq 0 ] ; then
	kubectl delete -f _out/manifests/release/kubevirt-cr.yaml && kubectl delete -f _out/manifests/release/kubevirt-operator.yaml
else
	echo "nothing to delete"
fi
