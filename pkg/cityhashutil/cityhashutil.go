package cityhashutil

import (
	//"fmt"
	"bytes"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"

	//uu "unicode/utf8"
	"github.com/JosephZoeller/cityhash"
)

// GetStrCode64Hash does
func GetStrCode64Hash(permutation string) uint64 {
	runePermute := []rune(permutation) // maybe something in bytes.runes for this
	encodedPermute := utf16.Encode(runePermute)

	sf := uint32(encodedPermute[0]) << 16
	sl := uint32(len(encodedPermute))
	var seed1 uint64 = uint64(sf + sl)

	input := append([]byte(permutation), 0)

	var mask uint64 = 0xFFFFFFFFFFFF

	return cityhash.CityHash64WithSeeds([]byte(input), 0x9ae16a3b2f90404f, seed1) & mask
}

// TestStrCode64Hash does
func TestStrCode64Hash() {
	collisiontests := []string{" ", "TENSION_NECK", "a", "joke", "hi", "how", "are", "you", "?"}

	for _, v := range collisiontests {
		f := GetStrCode64Hash(v)
		fmt.Printf("\n'%s' Post-CityHash: %x | %d\n\n", v, f, f)
	}
}

// DecodeUTF16 from https://gist.github.com/bradleypeabody/185b1d7ed6c0c2ab6cec
func DecodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}
