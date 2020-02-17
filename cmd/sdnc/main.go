package main

import (
	"fmt"

	"github.com/200106-uta-go/JKJP2/pkg/cityhashutil"
)

// SDN Controller Entry point.
// Listens for Client Query (hash)
// Broadcasts query to workers (hash)
// Listens for worker responses (hash + collision)
// Logs worker responses (hash + collision -> collisions.txt)
func main() {

	// rapid testing the cityhash algorithm with dismal results
	const knownCollision string = ("TENSION_NECK")

	f := cityhashutil.GetStrCode64Hash(knownCollision)
	fmt.Printf("\nPost-CityHash: %s\n", f)

	//go listenForClient()
	//go listenForWorker() currently, forwarding the client's request to the workers will then wait for the workers to return the collision as a response (0-5 seconds)
	/*
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT)
		<-signalChan
	*/
}
