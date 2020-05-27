package main

import (
	//"net/http"
	"fmt"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var SDNCAddr string
var collisionChan chan cityhashutil.ColliderResponse
var exitChan chan bool
var remainingHashes []uint64
var start time.Time

func main() {
	start = time.Now()
	collisionChan = make(chan cityhashutil.ColliderResponse)
	exitChan = make(chan bool)
	//http.HandleFunc("/SDNCToWorker", ListenForSDNC)

	//go http.ListenAndServe(":8080", nil)
	//go postCollisions()
	
	go postCollisionsTest()
	go AlgorithmTest()

	<-exitChan
	fmt.Println("END")
}
