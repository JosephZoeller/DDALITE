package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
	terra "github.com/JosephZoeller/DDALITE/pkg/terrautil"
)

var (
	overIps          []string
	instanceCount    *int
	clientAddr string
)

func init() {
	log.SetFlags(log.Llongfile)
	instanceCount = flag.Int("c", 1, "count - Determines how many instances to spin up. Default 1.")
	flag.Parse()
}

func main() {
	spinUp(*instanceCount)

	http.HandleFunc("/ClientToSDNC", listenForClient)
	http.HandleFunc("/WorkerToSDNC", listenForWorker)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		err := http.ListenAndServe(":666", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Kill
		}
	}()

	<-signalChan
	fmt.Println("Beginning Infrastructure Teardown...")
	
	tErr := kubeutil.TearDown()
	if tErr != nil {
		log.Printf(tErr.Error())
	}
	terra.TearDown()

	fmt.Println("Teardown complete. Rerun the SDNC to rebuild the infrastructure, or exit the ssh connection and tear down the master with 'make destroy_master'.")
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