package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Get service port value from environment variable.
	port := os.Getenv("SERV_PORT")

	// Response function initiator.
	http.HandleFunc("/", Resp)

	// Listen http via service port as definded in enviroment variable.
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
