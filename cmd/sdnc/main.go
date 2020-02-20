package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// SDN Controller Entry point.
// Listens for Client Query from Reverse Proxy (hash) Form: http://my.ip/:8080?hash=s0m3h4sh&instances=3
// Broadcasts query to workers (hash)
// Listens for worker responses (hash + collision)
// Logs worker responses (hash + collision -> collisions.txt)
func main() {

	fmt.Println("SDN Controller now listening on port 8080")
	http.Handle("/", http.FileServer(http.Dir("cmd/sdnc/html")))
	http.HandleFunc("/loading", processing)
	http.HandleFunc("/result", listenForClient)
	http.ListenAndServe(":8080", nil)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	<-signalChan
}
