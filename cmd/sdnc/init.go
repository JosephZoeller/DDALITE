package main

import (
	"os"
	"fmt"
)

var workerAddrs []string

func init() {
	for i := 0; i < 2; i++ {
		wrkAddr := os.Getenv(fmt.Sprintf("wrkAddr_%d", 1))
		workerAddrs = append(workerAddrs, wrkAddr)
	}
}