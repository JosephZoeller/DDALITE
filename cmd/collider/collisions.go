package main

import (
	//"fmt"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

// recursive function to build combinations of dictionary words and compare them with the hashes
func findCollisions(workParams cityhashutil.ColliderSpecifications) {
	remainingHashes = workParams.InputHashes

	for _, initWord := range workParams.Dictionary {
		checkCandidate(initWord, uint32(len(initWord)))
		combineRecurse(initWord, 1, &workParams)
	}
	for i := len(remainingHashes) - 1; i >= 0; i-- {
		collisionChan <- cityhashutil.ColliderResponse{Hashed: remainingHashes[i], Err: "[Hash Not Found (Search Complete)]"}
		remainingHashes = remainingHashes[:i]
	}
}

func combineRecurse(base string, fathometer int, workParams *cityhashutil.ColliderSpecifications) {
	for _, word := range workParams.Dictionary {
		candidate := base + workParams.Delimiter + word
		checkCandidate(candidate, uint32(len(candidate))) // computing the candidate's length once should reduce processing with a long list of input hashes

		if fathometer+1 < workParams.Depth {
			combineRecurse(candidate, fathometer+1, workParams)
		}
	}
}

func checkCandidate(candidate string, candidateLen uint32) {
	//fmt.Println(candidate) // for testing only! printing is a needless expense

	for i := 0; i < len(remainingHashes); i++ {
		if compare(candidate, remainingHashes[i], &candidateLen) { 
			collisionChan <- cityhashutil.ColliderResponse{Hashed: remainingHashes[i], Unhashed: candidate, Err: ""}
			remainingHashes = remove(remainingHashes, i)
			return // ASSUME: no duplicates
		}
	}
}

func compare(candidate string, hash uint64, candidateLen *uint32) bool {
	return cityhashutil.GetStrCode64HashWithLen(candidate, candidateLen) == hash
}

// from https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s []uint64, i int) []uint64 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
