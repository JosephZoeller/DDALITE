package main

import (
	"io/ioutil"
	"os"

	"testing"
)

func Test(t *testing.T) { // TENSION_NECK == 60db807ad6c6 (106495869638342) // apple == b10f405db64b (194679062509131) Apple == 4e1ec7e1511b (85894109417755)
	apbytes := []byte("Apple")
	er := ioutil.WriteFile("tmpDic.txt",apbytes, 0777)
	if er != nil {
		t.Error("Dictionary not created - ", er)
	}

	f, er := getDictionary("tmpDic.txt")
	if er != nil {
		t.Error("Could not retrieve dictionary")
	}
	found := findCollisionFile(f, "85894109417755", 0, 450000)
	if found == "Apple" {
		t.Log(found)
	} else {
		os.Remove("tmpDic.txt")
		t.Error("Known dictionary hash not found")
	}
	os.Remove("tmpDic.txt")
}

func Benchmark(b *testing.B) {
	apbytes := []byte("Apple")
	er := ioutil.WriteFile("tmpDic.txt",apbytes, 0777)
	if er != nil {
		b.Error("Dictionary not created - ", er)
	}

	f, er := getDictionary("tmpDic.txt")
	if er != nil {
		b.Error("Could not retrieve dictionary")
	}
	for i := 0; i < b.N; i++ {
		findCollisionFile(f, "1", 0, i)
	}
	os.Remove("tmpDic.txt")

}
/*
func TestDB(t *testing.T) { // Apple == 4e1ec7e1511b (85894109417755)
	var er error

	flyerDB, er = flyerdbutil.ConnectHard()
	if er != nil {
		t.Error("Connection Failed - ", er)
	}

	found := findDBCollision("85894109417755", 1, 20000)
	if found != "" {
		t.Log("Found: ", found)
	} else {
		t.Error("Known dictionary hash not found: 85894109417755")
	}
}

func TestGetWord(t *testing.T) {
	var word string
	var er error
	flyerDB, er = flyerdbutil.ConnectHard()
	if er != nil {
		t.Error("Connection Failed - ", er)
	}
	row := flyerDB.QueryRow("SELECT idx FROM dic WHERE word = $1", "zebra")
	er = row.Scan(&word)
	if er != nil {
		t.Error("Query Failed - ", er)
	}
}

func BenchmarkGetWord(b *testing.B) {
	var word string
	flyerDB, _ = flyerdbutil.ConnectHard()
	for i := 0; i < b.N; i++ {

		row := flyerDB.QueryRow("SELECT word FROM dic WHERE idx = $1", i)
		row.Scan(&word)
	}
}
*/