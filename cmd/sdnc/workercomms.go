package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"io"
	"log"
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var collision []cityhashutil.ColliderResponse

func listenForWorker(rw http.ResponseWriter, req *http.Request) {
	log.Println("Received data from worker...")
	collision := cityhashutil.ColliderResponse{}
	err := json.NewDecoder(req.Body).Decode(&collision)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully decoded worker data. Passing it onto client..")
		clientMsg := cityhashutil.MessageResponse{ Message: fmt.Sprint(collision.Hashed, " | ", collision.Unhashed)}
		post, _ := json.Marshal(clientMsg)
		postAddr := fmt.Sprintf("http://%s:8080/SDNCToClient", clientAddr)
		http.Post(postAddr, "application/json", bytes.NewReader(post))
	}

}

func sendToWorkers(workSpec cityhashutil.ClientSpecifications) {
	for i, addr := range overIps {
		log.Println("Sending work to: ", addr)
		if i < len(workSpec.Colliders) {
			work, _ := json.Marshal(workSpec.Colliders[i])
			postAddr := fmt.Sprintf("http://%s:8080/SDNCToWorker", addr)
			rsp, err := http.Post(postAddr, "application/json", bytes.NewReader(work))

			msg := cityhashutil.MessageResponse{}
			if err != nil {
				panic(err)
			} else {
				err = json.NewDecoder(rsp.Body).Decode(&msg)
				if err != nil {
					log.Println("Failed to decode server response - ", err)
				} else {
					log.Println(msg.Message)
				}
			}
		}
	}
}