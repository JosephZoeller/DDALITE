package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func ListenForSDNC(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request recieved from SDNC: ", req.RemoteAddr)
	SDNCAddr = strings.Split(req.RemoteAddr, ":")[0]

	work := cityhashutil.ColliderSpecifications{}
	msg := cityhashutil.MessageResponse{}

	err := json.NewDecoder(req.Body).Decode(&work)
	if err != nil {
		msg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(msg)
		log.Println("Failed to decode request from SDNC.")
		return
	}

	go findCollisions(work)

	msg.Message = "Worker is seeking collisions..."
	json.NewEncoder(rw).Encode(msg)
}

func postCollisions() {
	post, err := json.Marshal(<- collisionChan)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sending False Data...")
		postAddr := fmt.Sprintf("http://%s:666/WorkerToSDNC", SDNCAddr)
		http.Post(postAddr, "application/json", bytes.NewReader(post))
	}
}
