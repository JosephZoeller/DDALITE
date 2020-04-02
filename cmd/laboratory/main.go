package main

import (
	//"bufio"
	"errors"
	"fmt"

	//"os"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func main() {
	labSender()
}

func labSender() {
	clientBuiltFunc := func(inputs []string, cFunc func(inputHash, candidate string) bool) {
		zeroInput := inputs[0]
		generatedGuess := "Apple"
		if cFunc(zeroInput, generatedGuess) {
			labReturner(cityhashutil.HashOutParams{Hashed: zeroInput, Unhashed: generatedGuess, Err: nil})
		} else {
			labReturner(cityhashutil.HashOutParams{Hashed: zeroInput, Unhashed: "----------", Err: errors.New("Failed to find collision")})
		}

	}

	labExp := cityhashutil.HashInParams{
		InputHashes: []string{"85894109417755"},
		HashType:    "StrCode64",
		CompareFunc: clientBuiltFunc,
	}

	labReceiver(labExp)
}

func labReceiver(input cityhashutil.HashInParams) {
	var comparerFunc func(string, string) bool

	switch input.HashType {
	case "StrCode64":
		comparerFunc = func(inputHash, candidate string) bool {
			candidateHash := fmt.Sprintf("%d", cityhashutil.GetStrCode64Hash(candidate))
			return candidateHash == inputHash
		}
	default:
		return
	}

	input.CompareFunc(input.InputHashes, comparerFunc)
}

func labReturner(output cityhashutil.HashOutParams) {
	fmt.Println("temp collision: ", output.Unhashed, " | Input hash: ", output.Hashed)
	if output.Err != nil {
		fmt.Println(output.Err)
	}
}
