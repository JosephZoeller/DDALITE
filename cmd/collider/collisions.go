package main

import (
	"fmt"
	"sync"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var workParams cityhashutil.ColliderSpecifications
var wg sync.WaitGroup

// recursive function to build combinations of dictionary words and compare them with the hashes
func findCollisions(specs cityhashutil.ColliderSpecifications) {
	workParams = specs
	remainingHashes = workParams.InputHashes

	if workParams.ChipCount == 0 {
		candidateAugment := workParams.StartsWith + workParams.Delimiter + workParams.EndsWith
		checkCandidate(candidateAugment, uint32(len(candidateAugment)))
	}

	workParams.StartsWith = workParams.StartsWith + workParams.Delimiter
	workParams.EndsWith = workParams.Delimiter + workParams.EndsWith

	if workParams.ChipCount == 1 {
		for _, candidate := range workParams.Dictionary {
			candidateAugment := workParams.StartsWith + candidate + workParams.EndsWith
			checkCandidate(candidateAugment, uint32(len(candidateAugment)))
		}

	} else {
		for _, candidate := range workParams.Dictionary {
			wg.Add(1)
			go combineRecurseGoroutineLayer(candidate)
		}

	}

	debriefing()
}

func combineRecurseGoroutineLayer(candidate string) {
	combineRecurse(candidate, 1)
	wg.Done()
}

func combineRecurse(base string, fathometer int) {
	for _, word := range workParams.Dictionary {
		if fathometer+1 < workParams.ChipCount {
			combineRecurse(base+workParams.Delimiter+word, fathometer+1)
		} else {
			candidateAugment := workParams.StartsWith + base + workParams.Delimiter + word + workParams.EndsWith
			checkCandidate(candidateAugment, uint32(len(candidateAugment)))
		}
	}
}

func checkCandidate(candidate string, candidateLen uint32) { // string length is used in hashing, computing it beforehand should reduce processing with a long list of input hashes
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

func debriefing() {
	rHashCount := len(remainingHashes)
	collisionChan <- cityhashutil.ColliderResponse{Err: fmt.Sprintf("[Search Complete, %d hashes not found]", rHashCount)}
}
