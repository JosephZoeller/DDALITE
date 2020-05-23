package main

import (
	"log"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)


func AlgorithmTest() {

	work := cityhashutil.ColliderSpecifications{
		InputHashes: []uint64{0x4e1ec7e1511b,6285962488583}, //6285962488583
		Dictionary:  []string{"A", "p", "l", "e", "s"}, //, "s"
		Delimiter:   "",
		Depth:       15,
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