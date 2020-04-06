package main

import (
	"fmt"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func findCollisions(workParams cityhashutil.HashInParamsOnline, collissionChan chan cityhashutil.HashOutParams) {
	for i := 0; i < 5; i++ {
		postCollisions(cityhashutil.HashOutParams{Hashed: "1234", Unhashed: "TestUnhash", Err: ""})
		time.Sleep(time.Second * 10)
	}
}

func compare(inputHash string, candidate string) bool {
	candidateHash := fmt.Sprintf("%d", cityhashutil.GetStrCode64Hash(candidate))
	return candidateHash == inputHash
}
