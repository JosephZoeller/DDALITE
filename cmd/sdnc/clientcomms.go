package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
)

func listenForClient(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request recieved from client: ", req.RemoteAddr)
	clientAddr = strings.Split(req.RemoteAddr, ":")[0]

	workSpec := cityhashutil.ClientSpecifications{}

	srvMsg := cityhashutil.MessageResponse{}
	err := json.NewDecoder(req.Body).Decode(&workSpec)
	if err != nil {
		srvMsg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(srvMsg)
		log.Println("Failed to decode client Post.")
		return
	}
	dicCount := len(workSpec.Dictionaries)
	srvMsg.Message = "Successfully decoded client data... Sending work request to workers."
	
	kubeutil.SetUp(dicCount)
	log.Println("Checking pod readiness...")
	for !allPodsReady(dicCount) {
		log.Println("Rechecking...")
		time.Sleep(time.Second  * 5)
	}

	log.Println("Pods ready, gathering IPs...")
	refreshIps()

	json.NewEncoder(rw).Encode(srvMsg)
	sendToWorkers(workSpec)
}
