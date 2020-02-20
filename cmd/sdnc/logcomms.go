package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const collisionFileName = "collisions.txt"

// exportCollision logs the recieved hash-collision pair to a text file.
// Aside: this logic  would be better off in a sidecar

func exportCollision(hash, collision string) {
	var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	collisionFile, er := os.Create(collisionFileName)
	if er == nil {
		defer collisionFile.Close()
		logger = log.New(io.MultiWriter(collisionFile, os.Stdout), "", log.Ldate|log.Ltime)
	} else {
		logger.Println(fmt.Sprintf("[Log Manager]: Failed to create file for %s - %s", collisionFileName, er.Error()))
	}

	logger.Println("Hash: ", hash, " | Collision: ", collision)
}
