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
	http.HandleFunc("/SDNCToClient", listenForSDNC)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Interrupt
		}
	}()

	<-signalChan
}

func listenForSDNC(response http.ResponseWriter, request *http.Request) {
	m := cityhashutil.ResponseMessage{}

	err := json.NewDecoder(request.Body).Decode(&m)
	if err != nil {
		log.Println("Error: Failed to decode server message - ", err)
	} else {
		log.Println(m.Message)
	}
}
