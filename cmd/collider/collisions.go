package main

import (
	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

// recursive function to build combinations of dictionary words and compare them with the hashes
func findCollisions(workParams cityhashutil.ColliderSpecifications) {
	for _, initWord := range workParams.Dictionary {
		checkCandidate(initWord, workParams.InputHashes)
		combineRecurse(initWord, 1, workParams)
	}
}

func combineRecurse(base string, fathometer int, workParams cityhashutil.ColliderSpecifications) {
	for _, word := range workParams.Dictionary {
		candidate := base + workParams.Delimiter + word
		checkCandidate(candidate, workParams.InputHashes)

		if fathometer+1 < workParams.Depth {
			combineRecurse(candidate, fathometer+1, workParams)
		}
	}
}

func checkCandidate(candidate string, hashes []uint64) {
	//fmt.Println(candidate) // for testing only! printing is a needless expense

	for _, hash := range hashes {
		if compare(candidate, hash) {
			collisionChan <- cityhashutil.ColliderResponse{Hashed: hash, Unhashed: candidate, Err: ""}
		}
	}
}

func compare(candidate string, hash uint64) bool {
	return cityhashutil.GetStrCode64Hash(candidate) == hash
}
