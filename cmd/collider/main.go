package main

import (
	"fmt"
	"log"
	"net/http"
)

var SDNCAddr string

func main() {
	fmt.Printf("Running Collider Server on port 8080...")
	http.HandleFunc("/", ListenForSDNC)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
