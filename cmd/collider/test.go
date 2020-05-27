package main

import (
	//"fmt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func AlgorithmTest() {
	work := cityhashutil.ColliderSpecifications{}
	/*
		work = cityhashutil.ColliderSpecifications{
			InputHashes: []uint64{0x4e1ec7e1511b},
			Dictionary:  []string{"p", "l"},
			Delimiter:   "",
			StartsWith:  "A",
			EndsWith:    "e",
			Words:       3,
		}
		findCollisions(work)
	*/

	unmatchedHashesFilename := "resources/AppleHash.txt"
	chipDictionaryFilename := "resources/chipDictionary.txt"
	literalDictionaryFilename := "resources/Dictionary.txt"

	unmatchedHashesFile, Err := os.Open(unmatchedHashesFilename)
	if Err != nil {
		panic(Err)
	}
	scanner := bufio.NewScanner(unmatchedHashesFile)
	unknownHashesUint64 := make([]uint64, 0)
	for scanner.Scan() {
		meshParsed, Err := strconv.ParseUint(scanner.Text(), 10, 64)
		if Err != nil {
			panic(Err)
		}
		unknownHashesUint64 = append(unknownHashesUint64, meshParsed)
	}

	chipDictionaryFile, Err := os.Open(chipDictionaryFilename)
	if Err != nil {
		panic(Err)
	}
	scanner = bufio.NewScanner(chipDictionaryFile)
	chipDictionary := make([]string, 0)
	for scanner.Scan() {
		chipDictionary = append(chipDictionary, scanner.Text())
	}
	//chipDictionary = []string{""}
	//chipDictionary = []string{"_","e", "t", "a", "o", "i", "n", "s", "r", "h", "l", "d", "c", "u", "m", "f", "p", "g", "w", "y", "b", "v", "k", "x", "j", "q", "z",
	//							"", "E", "T", "A", "O", "I", "N", "S", "R", "H", "L", "D", "C", "U", "M", "F", "P", "G", "W", "Y", "B", "V", "K", "X", "J", "Q", "Z"}
	//chipDictionary = []string{"E", "T", "A", "O", "I", "N", "S", "R", "H", "L", "D", "C", "U", "M", "F", "P", "G", "W", "Y", "B", "V", "K", "X", "J", "Q", "Z", "_"}
	//chipDictionary = []string{"A","e", "t", "a", "o", "i", "n", "s", "r", "h", "l", "d", "c", "u", "m", "f", "p", "g", "w", "y", "b", "v", "k", "x", "j", "q", "z", "_"}
	//chipDictionary = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	literalDictionaryFile, Err := os.Open(literalDictionaryFilename)
	if Err != nil {
		panic(Err)
	}
	scanner = bufio.NewScanner(literalDictionaryFile)
	literalDictionary := make([]string, 0)
	for scanner.Scan() {
		literalDictionary = append(literalDictionary, scanner.Text())
	}

	//var wmuUnknownMatInstances = []uint64{179744920977897,80962542339984}
	//var wmuUnknownMeshGroups = []uint64{0xb7e8a16adbbd, 0x56a033934dc}

	fmt.Println("Files loaded. Processing...")
	for _, dicWord := range literalDictionary {
		casing := strings.Title(dicWord)
		work = cityhashutil.ColliderSpecifications{
			InputHashes: unknownHashesUint64,
			Dictionary:  chipDictionary,
			Delimiter:   "",
			ChipCount:   0,
			StartsWith:  fmt.Sprintf("%s", casing),
			EndsWith:    "",
		}
		//fmt.Println(casing)
		findCollisions(work)
	}
	exitChan <- true
	/*
		for i := 0; i < 1000; i++ {

			work = cityhashutil.ColliderSpecifications{
				InputHashes: unknownHashesUint64,
				Dictionary: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "_"},
				Delimiter:  "",
				ChipCount:      5,
				StartsWith: fmt.Sprintf("SKL_%03d_", i),
				EndsWith:   "",
			}
			findCollisions(work)
		}
	*/
}

func postCollisionsTest() {
	f, _ := os.Create("logfile.log")
	//lognterminal := io.MultiWriter(f, os.Stdout)
	log.SetOutput(f)
	for {
		col := <-collisionChan
		log.Println(col)
		log.Println(time.Since(start))
		if len(remainingHashes) == 0 {
			log.Printf("Completed: %s", time.Since(start).String())
			exitChan <- true
		}
	}
}
