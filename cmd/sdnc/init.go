package main

import (
	"fmt"
	"os"
)

var (
	dictionaryLength string = 10 //!! WARNING THIS IS ONLY TEMPORARY PLEASE ADD INIT LOGIC TO GET ACTUAL LENGTH FROM DB
	workerAddrs      []string
)

// What is this for?

func init() {
	for i := 0; i < 2; i++ {
		wrkAddr := os.Getenv(fmt.Sprintf("wrkAddr_%d", 1))
		workerAddrs = append(workerAddrs, wrkAddr)
	}
}
