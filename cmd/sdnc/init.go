package main

import (
	"log"

	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
	"github.com/JosephZoeller/DDALITE/pkg/terra"
)

// Define ports for the different components that SDNC will be communicating with.
const (
	ColliderPort = "8080" // This is the expected port that the colliders will be listening on.
)

var (
	dictionaryLength int64 = 466550 //!! WARNING THIS IS ONLY TEMPORARY PLEASE ADD INIT LOGIC TO GET ACTUAL LENGTH FROM DB
	overIps []string
	instanceCount int = 1
)

func init() {
	log.SetFlags(log.Llongfile)
	spinUp(instanceCount)
}

func spinUp(iCnt int) {

	if iCnt <= 0 {
		log.Fatalf("Error instances=%d were not set.", iCnt)
	}

	// Initiate Terraform script to create EC2 instances.
	// NO LONGER RETURNS IPS. Use the ips here to log the EC2 underlay ips for safekeeping.
	terra.Provision(iCnt)

	// Launch deployment yaml in /kubernetes/deployment.yaml and return the pod private overlay ips.
	setUpErr := kubeutil.SetUp(iCnt)
	if setUpErr != nil {
		log.Fatalf(setUpErr.Error())
	}

	// Get Overlay IPs from current set of pods.
	overIps = make([]string, 0)
	myPods := kubeutil.PodInfo()

	for _, v := range myPods {
		overIps = append(overIps, v.IPaddr)
	}
}
