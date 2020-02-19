package main

import (
	"database/sql"
	"log"

	"github.com/200106-uta-go/JKJP2/pkg/flyerdbutil"
)

var flyerDB *sql.DB

func findDBCollision(inputHash string, startIndex, searchLength int) string {
	// Connect to database
	/*
		dbConfig, er := flyerdbutil.ReadConfig(dbConfigFilePath)
		if er != nil {
			log.Fatal("Failed to get Database Config data - ", er.Error())
		}
	*/
	var er error
	flyerDB, er = flyerdbutil.ConnectHard()
	if er != nil {
		log.Fatal("Database connection failed - ", er.Error())
	}
	defer flyerDB.Close()

	// iterate over database strings
	j := startIndex + searchLength
	for i := startIndex; i < j; i++ {
		word, er := getWordFromDB(i)
		if er != nil {
			return "" // get ready to redo this
		}

		if compare(inputHash, word) {
			return word
		}
	}
	return ""
}

func getWordFromDB(id int) (string, error) { // query database for a word
	var word string

	row := flyerDB.QueryRow("SELECT word FROM dic WHERE idx = $1", id)
	er := row.Scan(&word)

	return word, er
}
