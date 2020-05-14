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

func refreshIps() {
	myPods := kubeutil.PodInfo()

	overIps = make([]string, 0)
	for _, v := range myPods {
		log.Println(v.IPaddr)
		overIps = append(overIps, v.IPaddr)
	}
}
