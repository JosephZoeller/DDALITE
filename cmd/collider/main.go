package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/200106-uta-go/JKJP2/cmd/collider/proc"
)

func main() {
	// Response function initiator.
	http.HandleFunc("/", proc.Resp)

	// Listen http port 8080.
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
