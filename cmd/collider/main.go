package main

import (
	"fmt"
	"log"
	"net/http"
)

var inPort int
var dbConfigFilePath string
var dictionaryFilePath string

func init() {
	inPort = 8080
	dbConfigFilePath = "config.yaml"
	dictionaryFilePath = "dictionary.txt"
}

func main() {

	// Response function initiator.
	http.HandleFunc("/", Resp)

	// Listen http via service port as definded in enviroment variable.
	fmt.Printf("Collider is running on port %d\n", inPort)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
