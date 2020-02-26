package main

import (
	"log"
	"net/http"
)

func listenForClient(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
	}

	hash := req.FormValue("hash")
	if hash == "" {
		http.Error(rw, "Hash parse not found", http.StatusInternalServerError)
	}

	resp := sendToWorkers(hash, overIps)
	if resp != "" {
		log.Printf("Worker Returned Collision: %s\n", resp)
		exportCollision(hash, resp)
	}
}
