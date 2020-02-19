package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/200106-uta-go/JKJP2/pkg/cityhashutil"
	"github.com/200106-uta-go/JKJP2/pkg/kubeutil"
	"github.com/200106-uta-go/JKJP2/pkg/terra"
)

// listenForClient awaits a query (curl request) from the client. Upon recieving a request, the hash is handed out to the worker addresses.
func listenForClient() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Parse Form sent from reverse proxy. Set variables equal and check they aren't empty.
		err := req.ParseForm()
		if err != nil {
			http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
		}

		// Double check that variables actually have a value. Send a bad request
		// status if client is derp.
		hash := req.FormValue("hash")
		instanceCount := req.FormValue("instances")

		if hash == "" || instanceCount == "" {
			http.Error(rw, fmt.Sprintf("Error hash=%s and instances=%s were not set.", hash, instanceCount), http.StatusBadRequest)
			return
		}

		// Initiate Terraform script to create EC2 instances.
		// NO LONGER RETURNS IPS. Use the ips here to log the EC2 underlay ips for safekeeping.
		terra.Provision(instanceCount)

		// Launch deployment yaml in /kubernetes/deployment.yaml and return the pod private overlay ips.
		setUpErr := kubeutil.SetUp(instanceCount)
		if setUpErr != nil {
			http.Error(rw, setUpErr.Error(), http.StatusInternalServerError)
			return
		}

		// Get Overlay IPs from current set of pods.
		overIps := make([]string, 0)
		myPods := kubeutil.PodInfo()

		for _, v := range myPods {
			overIps = append(overIps, v.IPaddr)
		}

		// Send hash to each pod at each overlay ip.
		resp := sendToWorkers(hash, overIps)
		defer resp.Body.Close()

		// Store hash and collision in struct
		var myCollision cityhashutil.HashCollision

		jsErr := json.NewDecoder(resp.Body).Decode(&myCollision)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error decoding json: Error == %v", jsErr), http.StatusInternalServerError)
			kubeutil.TearDown()
			return
		}

		// Log result to stdoutput. May want to route logs to different location later.
		log.Printf("Worker Returned Collision: %v\n", myCollision)
		exportCollision(myCollision.InputHash, myCollision.Collision)

		// Wrap up myCollision into json because you do not want to read response body multiple times.
		js, err := json.Marshal(myCollision)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error marshaling json: Error == %v", err), http.StatusInternalServerError)
			kubeutil.TearDown()
			return
		}

		// Just pass on body from worker back to reverse proxy after marshaling.
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(js)

		// Tear down kubernetes pods and then ec2 instances to save money.
		tErr := kubeutil.TearDown()
		if tErr != nil {
			log.Printf(tErr.Error())
		}

	})

	http.ListenAndServe(":8080", nil)
}
