package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func getDictionary(textfile string) (*os.File, error) {

	return os.Open(textfile)

}

func findCollisionFile(dictionary *os.File, inputHash string, startIndex, searchLength int) string {

	sc := bufio.NewScanner(dictionary)

	// iterate over file strings
	j := startIndex + searchLength
	for i := 0; i < startIndex; i++ {
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
