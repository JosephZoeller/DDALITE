package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/200106-uta-go/JKJP2/pkg/cityhashutil"
	"github.com/200106-uta-go/JKJP2/pkg/flyerdbutil"
	_ "github.com/lib/pq"
)

var flyerDB *sql.DB

func findCollision(inputHash string, startIndex, searchLength int) string {
	// Connect to database
	dbConfig, er := flyerdbutil.ReadConfig(dbConfigFilePath)
	if er != nil {
		log.Fatal("Failed to get Database Config data - ", er.Error())
	}
	flyerDB, er = flyerdbutil.Connect("server", dbConfig)
	if er != nil {
		log.Fatal("Database connection failed - ", er.Error())
	}
	defer flyerDB.Close()

	// iterate over database strings
	j := startIndex + searchLength
	for i := startIndex; i < j; i++ {
		word, er := getWordFromDictionary(i)
		if er != nil {
			return "" // get ready to redo this
		}
		compareHash := fmt.Sprintf("%d", cityhashutil.GetStrCode64Hash(word))

		if compareHash == inputHash {
			return word
		}
	}
	return ""
}

func getWordFromDictionary(id int) (string, error) { // query database for a word
	var word string

	row := flyerDB.QueryRow("SELECT word FROM dic WHERE idx = $1", id)
	er := row.Scan(&word)

	return word, er
}
