package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)
/*
func listenForClientLegacy(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
		return
	}

	hash := req.FormValue("hash")
	if hash == "" {
		http.Error(rw, "Hash parse not found", http.StatusInternalServerError)
		return
	}

	resp := sendToWorkersLegacy(hash)
	if resp != "" {
		log.Printf("Worker Returned Collision: %s\n", resp)
		exportCollision(hash, resp)
	}
}
*/

func listenForClient(rw http.ResponseWriter, req *http.Request) {
	workSpec := cityhashutil.ClientPost{}

	srvMsg := cityhashutil.ResponseMessage{}
	err := json.NewDecoder(req.Body).Decode(&workSpec)
	if err != nil {
		srvMsg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(srvMsg)
		log.Println("Failed to decode client Post.")
	} else {
		srvMsg.Message = "Successful decode"
		json.NewEncoder(rw).Encode(srvMsg)
	}

	sendToWorkers(workSpec)
}
