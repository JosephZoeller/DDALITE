package main

import (
	"log"
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
	terra "github.com/JosephZoeller/DDALITE/pkg/terrautil"
)

func fakeSpinUp() {
	log.Printf("\nFaking Spinup...\n")
	overIps = make([]string, 0)
	overIps = append(overIps, "localhost")
}

func spinUp(iCnt int) {

	if iCnt <= 0 {
		log.Fatalf("Error instances=%d were not set.", iCnt)
	}

	// Initiate Terraform script to create EC2 instances.
	// NO LONGER RETURNS IPS. Use the ips here to log the EC2 underlay ips for safekeeping.

	log.Printf("\nProvisioning %d instances with Terraform...\n", iCnt)
	terra.Provision(iCnt)
	log.Println("Provisioning complete. Building Kubernetes environment...")

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
	log.Println("Kubernetes Environment built.")
}

func startTeardown(rw http.ResponseWriter, req *http.Request) {
	log.Println("Beginning Infrastructure Teardown...")

	tErr := kubeutil.TearDown()
	if tErr != nil {
		log.Printf(tErr.Error())
	}
	terra.TearDown()

	log.Println("Teardown complete. Rerun the SDNC to rebuild the infrastructure, or exit the ssh connection and tear down the master with 'make destroy_master'.")
}