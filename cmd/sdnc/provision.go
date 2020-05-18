package main

import (
	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
)

func allPodsReady(count int) bool {
	myPods := kubeutil.PodInfo()

	if len(myPods) == count {
		for _, v := range myPods {
			if v.Status != "Running" {
				return false
			}
		}
		return true
	}
	
	return false
}
