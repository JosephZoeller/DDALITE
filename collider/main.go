package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var inPort string
var dbConfigFilePath string

func init() {
	inPort = os.Getenv("SERV_PORT")
	dbConfigFilePath = "config.yaml"
}

func main() {
	// Response function initiator.
	http.HandleFunc("/", Resp)

	// Listen http via service port as definded in enviroment variable.
	fmt.Printf("Collider is running on port %s\n", inPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", inPort), nil))
}
