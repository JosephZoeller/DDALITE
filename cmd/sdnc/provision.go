package main

import (
	"log"

	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
)

func fakeIps() {
	log.Printf("\nFaking Spinup...\n")
	overIps = make([]string, 0)
	overIps = append(overIps, "localhost")
}

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

func refreshIps() {
	myPods := kubeutil.NodeInfo()

	overIps = make([]string, 0)
	for _, v := range myPods {
		log.Println("Pod IP: " + v.InternalIP)
		overIps = append(overIps, v.InternalIP)
	}
}
