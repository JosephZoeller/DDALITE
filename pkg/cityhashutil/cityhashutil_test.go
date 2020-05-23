package cityhashutil

import (
	"testing"
)

func Test(t *testing.T) {

	collisiontests := []string{"Apple", " ", "TENSION_NECK", "a", "joke", "hi", "how", "are", "you", "?"}

	for _, v := range collisiontests {
		f := GetStrCode64Hash(v)
		t.Logf("\n'%s' Post-CityHash: %x | %d\n\n", v, f, f)
	}
}