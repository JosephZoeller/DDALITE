package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	overIps    []string
	setup      = false
	clientAddr string
)

func main() {
	http.HandleFunc("/ClientToSDNC", listenForClient)
	http.HandleFunc("/WorkerToSDNC", listenForWorker)
	http.HandleFunc("/teardown", startTeardown)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		err := http.ListenAndServe(":666", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Kill
		}
	}()

	<-signalChan
}
