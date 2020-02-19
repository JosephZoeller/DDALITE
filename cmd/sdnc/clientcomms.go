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
		hash := req.FormValue("hash").Form.Get("hash")
		instanceCount := req.FormValue("instances")

		if hash == "" || instanceCount == "" {
			http.Error(rw, fmt.Sprintf("Error hash=%s and instances=%s were not set.", hash, instanceCount), http.StatusBadRequest)
			return
		}

		// Initiate Terraform script to create EC2 instances.
		// Use the ips here to log the EC2 underlay ips for safekeeping.
		ips := terra.Provision(instanceCount)
		fmt.Printf("ECS have been created @ %v", ips)

		// Launch deployment yaml in /kubernetes/deployment.yaml and return the pod private overlay ips.
		overIps := kubeutil.SetUp(instanceCount)

		// Send hash to each pod at each overlay ip.
		resp := sendToWorkers(hash, overIps)
		resp.Body.Close()

		// Store hash and collision in struct
		var myCollision cityhashutil.HashCollision

		err := json.NewDecoder(resp.Body).Decode(&myCollision)
		if err != nil {
			http.Error(rw, fmt.Printf("Error decoding json: Error == %v", err), http.StatusInternalServerError)
			return
		}

		// Log result to stdoutput. May want to route logs to different location later.
		log.Printf("Worker Returned Collision: %v\n", myCollision)

		// Wrap up myCollision into json because you do not want to read response body multiple times.
		js, err := json.Marshal(myCollision)
		if err != nil {
			http.Error(rw, fmt.Printf("Error marshaling json: Error == %v", err), http.StatusInternalServerError)
			return
		}

		// Just pass on body from worker back to reverse proxy after marshaling.
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(js)
	})

	http.ListenAndServe(":8080", nil)
}
