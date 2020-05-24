package main

import (
	"log"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)


func AlgorithmTest() {

	work := cityhashutil.ColliderSpecifications{
		InputHashes: []uint64{0x4e1ec7e1511b,6285962488583}, //6285962488583
		Dictionary:  []string{"p", "l", "e"}, //, "s"
		Delimiter:   "",
		Contains: "",
		StartsWith: "A",
		EndsWith: "s",
		Depth:       5,
	}

	findCollisions(work)
}

func postCollisionsTest() {
	for {
		log.Println(<-collisionChan)
		log.Println(time.Since(start))
		if len(remainingHashes) == 0 {
			exitChan <- true
		}
	}
}