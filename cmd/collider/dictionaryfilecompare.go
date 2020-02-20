package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/200106-uta-go/JKJP2/pkg/cityhashutil"
)


func findCollisionFile(inputHash string, startIndex, searchLength int) string {
	
	dictionary, _ := os.Open(dictionaryFilePath)
	defer dictionary.Close()
	sc := bufio.NewScanner(dictionary)
	
	// iterate over file strings
	j := startIndex + searchLength
	for i:= 0; i < startIndex; i++ {
		t := sc.Scan()
		if !t {
			return ""
		}
	}
	for i := startIndex; i < j; i++ {
		t := sc.Scan()
		if !t {
			return ""
		}

		if compare(inputHash, sc.Text()) {
			return sc.Text()
		}
	}
	return ""
}

func compare(inputHash string, candidate string) bool {
	candidateHash := fmt.Sprintf("%d", cityhashutil.GetStrCode64Hash(candidate))
	return candidateHash == inputHash
}