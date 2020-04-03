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
	http.HandleFunc("/sdnc", listenForSDNC)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	
	go func() {
		err := http.ListenAndServe("localhost:999", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Kill
		}
	}()

	<- signalChan
}

func listenForSDNC(response http.ResponseWriter, request *http.Request) {
	m := cityhashutil.MessageToClient{}
	
	err := json.NewDecoder(request.Body).Decode(&m)
	if err != nil {
		log.Println("Error: Failed to decode server message - ", err)
	} else {
		log.Println(m.Message)
	}

}