package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/200106-uta-go/JKJP2/pkg/terra"
)

// listenForClient awaits a query (curl request) from the client. Upon recieving a request, the hash is handed out to the worker addresses.
func listenForClient() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Parse Form sent from reverse proxy. Set variables equal and check they aren't empty.
		err := req.ParseForm()
		if err != nil {
			log.Panic("Error Parsing Client Request Form")
		}

		// Double check that variables actually have a value. Send a bad request
		// status if client is derp.
		hash := req.FormValue("hash").Form.Get("hash")
		instanceCount := req.FormValue("instances")

		if hash == "" || instanceCount == "" {
			rw.WriteHeader(http.StatusBadRequest)
			errMsg := fmt.Sprintf("Error hash=%s and instances=%s were not set.", hash, instanceCount)
			rw.Write([]byte(errMsg))
			break
		}

		// Initiate Terraform script to create EC2 instances.
		ips := terra.Provision(instanceCount)

		// Launch deployment yaml in /kubernetes/deployment.yaml
		kubeutil.SetUp()

		sendToWorkers(hash, ips)
	})

	http.ListenAndServe(":8080", nil)
}
