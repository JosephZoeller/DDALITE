package cityhashutil

import (
	//"fmt"
	"fmt"
	u "unicode/utf16"

	//uu "unicode/utf8"
	"github.com/Upliner/cityhash"
)

// GetStrCode64Hash does
func GetStrCode64Hash(permutation string) string {
	runePermute := []rune(permutation) // bytes.runes
	encodedPermute := u.Encode(runePermute)
	var sf uint64 = uint64(encodedPermute[0]) << 16
	var sl uint64 = uint64(len(encodedPermute))

	var seed1 uint64 = (sf + sl)
	fmt.Printf("Seed0: %d, Seed1: %d\n", uint64(0x9ae16a3b2f90404f), seed1) // known good
	encodedPermute = append(encodedPermute)
	t := string(u.Decode(encodedPermute))
	fmt.Printf("permutation input: %d", append([]byte(t), 0))
	var mask uint64 = 0xFFFFFFFFFFFF

	return fmt.Sprintf("%x", cityhash.CityHash64WithSeeds([]byte(""), 0x9ae16a3b2f90404f, seed1)&mask)
	//known good! in C#, equivalent to (CityHash.CityHash.CityHash64WithSeeds(Encoding.ASCII.GetBytes(""), seed0, seed1) & mask).ToString("x");
}
