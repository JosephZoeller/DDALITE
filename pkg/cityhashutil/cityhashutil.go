package cityhashutil

import (
	"github.com/JosephZoeller/cityhash"
)

func GetStrCode64Hash(permutation string) uint64 {

	sf := uint32(permutation[0]) << 16
	sl := uint32(len(permutation))
	var seed1 uint64 = uint64(sf + sl)

	input := append([]byte(permutation), 0)

	var mask uint64 = 0xFFFFFFFFFFFF

	return cityhash.CityHash64WithSeeds([]byte(input), 0x9ae16a3b2f90404f, seed1) & mask
}

func GetStrCode64HashWithLen(permutation string, sl *uint32) uint64 {

	sf := uint32(permutation[0]) << 16
	var seed1 uint64 = uint64(sf + *sl)

	input := append([]byte(permutation), 0)

	var mask uint64 = 0xFFFFFFFFFFFF

	return cityhash.CityHash64WithSeeds([]byte(input), 0x9ae16a3b2f90404f, seed1) & mask
}