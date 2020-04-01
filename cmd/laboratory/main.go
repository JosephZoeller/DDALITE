package main

import (
	//"bufio"
	"errors"
	"fmt"

	//"os"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func main() {
	lab()
	/*
		fmt.Println("HELLO THERE.")
		cmdreader := bufio.NewReader(os.Stdin)
		in, _ := cmdreader.ReadString('\n')
		if in != "GENERAL KENOBI\n" {
			fmt.Println("BYE THERE.")
		} else {
			fmt.Println("SO UNCIVILIZED")
		}
		fmt.Println("'", in, "'")
	*/
}

func lab() {
	testFunc := func(inputs []string, cFunc func(strVar string) bool) error {
		zeroInput := inputs[0]
		if cFunc(zeroInput) {
			return nil
		} else {
			return errors.New("cFunc didn't like your '" + zeroInput + "'.")
		}
	}

	exCompare := func(strVar string) bool {
		return strVar == "test"
	}

	labExp := cityhashutil.HashInParams{
		InputHashes: []string{"ok boomer"},
		HashType:    "StrCode64",
		CompareFunc: testFunc,
	}

	err := labExp.CompareFunc(labExp.InputHashes, exCompare)
	if err == nil {
		fmt.Println("Successful")
	} else {
		fmt.Println(err)
	}
}
