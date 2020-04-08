package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func listenForClient(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request recieved from client: ", req.RemoteAddr)
	clientAddr = strings.Split(req.RemoteAddr, ":")[0]

	workSpec := cityhashutil.ClientPost{}

	srvMsg := cityhashutil.ResponseMessage{}
	err := json.NewDecoder(req.Body).Decode(&workSpec)
	if err != nil {
		srvMsg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(srvMsg)
		log.Println("Failed to decode client Post.")
		return
	}
	srvMsg.Message = "Successfully decoded client data... "

	if !setup {
		srvMsg.Message += "Spinning up workers, this will take a moment... "
		spinUp(len(workSpec.Dictionaries))
		setup = true
	}
	srvMsg.Message += "Sending work request to workers."

	json.NewEncoder(rw).Encode(srvMsg)
	sendToWorkers(workSpec)
}
