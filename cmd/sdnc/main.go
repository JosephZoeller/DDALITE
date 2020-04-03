package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephZoeller/DDALITE/pkg/kubeutil"
	terra "github.com/JosephZoeller/DDALITE/pkg/terrautil"
)

// SDN Controller Entry point.
// Listens for Client Query from Reverse Proxy (hash) Form: http://my.ip/:8080?hash=s0m3h4sh&instances=3
// Broadcasts query to workers (hash)
// Listens for worker responses (hash + collision)
// Logs worker responses (hash + collision -> collisions.txt)
func main() {
	fmt.Println("SDN Controller now listening on port 8080")
	http.HandleFunc("/client", listenForClient)
	http.HandleFunc("/worker", listenForWorker)
	go http.ListenAndServe(":8080", nil)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	<-signalChan

	fmt.Println("Beginning Infrastructure Teardown...")
	// Tear down kubernetes pods and then ec2 instances to save money.
	tErr := kubeutil.TearDown()
	if tErr != nil {
		log.Printf(tErr.Error())
	}
	terra.TearDown()

	fmt.Println("Teardown complete. Rerun the SDNC to rebuild the infrastructure, or exit the ssh connection and tear down the master with 'make destroy_master'.")
}
