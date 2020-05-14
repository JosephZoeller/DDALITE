package main

import (
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var SDNCAddr string
var collisionChan chan cityhashutil.ColliderResponse

func main() {
	collisionChan = make(chan cityhashutil.ColliderResponse, 5)
	http.HandleFunc("/SDNCToWorker", ListenForSDNC)

	go http.ListenAndServe(":8080", nil)
	go postCollisions()
	
	/*
	http.HandleFunc("/Test", AlgorithmTest)
	go http.ListenAndServe("localhost:8080", nil)
	go postCollisionsTest()
	*/

	for {
	}
}
