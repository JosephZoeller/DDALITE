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
	fmt.Println("Received data from worker...")
	collision := cityhashutil.ColliderResponse{}
	err := json.NewDecoder(req.Body).Decode(&collision)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully decoded worker data. Passing it onto client..")
		clientMsg := cityhashutil.MessageResponse{Message: fmt.Sprint(collision.Hashed, " | ", collision.Unhashed)}
		post, _ := json.Marshal(clientMsg)
		postAddr := fmt.Sprintf("http://%s:8080/SDNCToClient", clientAddr)
		http.Post(postAddr, "application/json", bytes.NewReader(post))
	}

}

func sendToWorkers(workSpec cityhashutil.ClientSpecifications) {
	dicLen := len(workSpec.Dictionaries)
	for i := 0; i < dicLen; i++ {
		log.Println("Sending work to load balancer: ", ingressIp)
		work, _ := json.Marshal(cityhashutil.ColliderSpecifications{
			InputHashes: workSpec.InputHashes,
			Dictionary:  workSpec.Dictionaries[i],
			Delimiter:   workSpec.Delimiter,
			Depth:       workSpec.Depth,
		})

		msg := cityhashutil.MessageResponse{}
		postAddr := fmt.Sprintf("http://%s:8080/SDNCToWorker", ingressIp)
		rsp, err := http.Post(postAddr, "application/json", bytes.NewReader(work))
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
