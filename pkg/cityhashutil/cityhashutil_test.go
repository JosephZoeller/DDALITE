package cityhashutil

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	collisiontests := []string{" ", "TENSION_NECK", "a", "joke", "hi", "how", "are", "you", "?"}

	for _, v := range collisiontests {
		f := GetStrCode64Hash(v)
		fmt.Printf("\n'%s' Post-CityHash: %x | %d\n\n", v, f, f)
	}
}