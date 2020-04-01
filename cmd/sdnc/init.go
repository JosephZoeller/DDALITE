package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
	terra "github.com/JosephZoeller/DDALITE/pkg/terrautil"
)

// Define ports for the different components that SDNC will be communicating with.
const (
	ColliderPort = "8080" // This is the expected port that the colliders will be listening on.
)

var (
	dictionaryLength int64 = 466550 //!! WARNING THIS IS ONLY TEMPORARY PLEASE ADD INIT LOGIC TO GET ACTUAL LENGTH FROM DB
	overIps          []string
	instanceCount    *int
)

func init() {
	log.SetFlags(log.Llongfile)
	instanceCount = flag.Int("c", 1, "count - Determines how many instances to spin up. Default 1.")
	flag.Parse()
	spinUp(*instanceCount)
}

func spinUp(iCnt int) {

	if iCnt <= 0 {
		log.Fatalf("Error instances=%d were not set.", iCnt)
	}

	// Initiate Terraform script to create EC2 instances.
	// NO LONGER RETURNS IPS. Use the ips here to log the EC2 underlay ips for safekeeping.

	fmt.Printf("\nProvisioning %d instances with Terraform...\n", iCnt)
	terra.Provision(iCnt)
	fmt.Println("Provisioning complete. Building Kubernetes environment...")

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
	fmt.Println("Kubernetes Environment built.")
}
