package cityhashutil

import (
	"fmt"

	"github.com/Upliner/cityhash"
)
// GetStrCode64HashT1 does
func GetStrCode64HashT1(permutation []byte) uint64 {

	fmt.Println(uint64(int32(permutation[0]) << 16))
	fmt.Println(uint64(len(permutation) / 2))
	var sf uint64 = uint64(int32(permutation[0]) << 16)
	var sl uint64 = uint64(len(permutation) / 2)

	var seed1 uint64 = (sf + sl)
	t := append(permutation, 0x00, 0x00)
	fmt.Println(t)
	fmt.Println(seed1)

	// There's something really weird going on with CityHash64WithSeeds that I can't put my finger on.
	// I even opened up the original C program and both this cityhash library and the original gave me the same outputs,
	// but when I put the exact same bytes into my old C# program, it gives me something completely different
	// I know my original c# program works as intended and returns the appropriate hash, but this program and the original C program both seem to 
	// give me some unexpected number. I know that TENSION_HEAD maps to "60db807ad6c6", because that's the collision/hash pair directly from the game
	// however, this is returning "10016ce8bcf8". There's got to be a relationship between these numbers somehow

	// I'm almost certain that it has to do with C# using UTF16 chars and golang using UTF8 bytes to represent strings, but the solution just isn't clicking with me.

	return cityhash.CityHash64WithSeeds(t, 0x9ae16a3b2f90404f, seed1) & 0xFFFFFFFFFFFF

}