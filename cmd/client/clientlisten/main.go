package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func main() {
	f, _ := os.Create("logfile.log")
	log.SetOutput(f)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	http.HandleFunc("/SDNCToClient", listenForSDNC)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Interrupt
		}
	}()

	/*
	http.HandleFunc("/LocalToClient", listenForLocal)
	go func() {
		err := http.ListenAndServe("localhost:8081", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Interrupt
		}
	}()
	*/
	<-signalChan
}

func listenForSDNC(response http.ResponseWriter, request *http.Request) {
	m := cityhashutil.MessageResponse{}

	err := json.NewDecoder(request.Body).Decode(&m)
	if err != nil {
		log.Println("Error: Failed to decode server message - ", err)
	} else {
		log.Println(m.Message)
	}
}

func listenForLocal(response http.ResponseWriter, request *http.Request) {
	m := cityhashutil.ColliderResponse{}

	err := json.NewDecoder(request.Body).Decode(&m)
	if err != nil {
		log.Println("Error: Failed to decode worker message - ", err)
	} else {
		log.Println(m.Hashed, m.Unhashed, m.Err)
	}
}