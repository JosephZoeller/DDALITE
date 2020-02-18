package main

import (
	"fmt"
	"os"
)

var workerAddrs []string

// What is this for?

func init() {
	for i := 0; i < 2; i++ {
		wrkAddr := os.Getenv(fmt.Sprintf("wrkAddr_%d", 1))
		workerAddrs = append(workerAddrs, wrkAddr)
	}
}
