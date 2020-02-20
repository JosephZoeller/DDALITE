package main

import (
	"testing"

	"github.com/200106-uta-go/JKJP2/pkg/flyerdbutil"
)

func Test(t *testing.T) { // TENSION_NECK == 60db807ad6c6 (106495869638342) // apple == b10f405db64b (194679062509131) Apple == 4e1ec7e1511b (85894109417755)
	found := findCollisionFile("85894109417755", 0, 450000)
	if found != "" {
		t.Log(found)
	} else {
		t.Error("Known dictionary hash not found: 85894109417755")
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findCollisionFile("1", 0, i)
	}

}

func TestDB(t *testing.T) { // Apple == 4e1ec7e1511b (85894109417755)
	var er error

	flyerDB, er = flyerdbutil.ConnectHard()
	if er != nil {
		t.Error("Connection Failed - ", er)
	}

	found := findDBCollision("36478381552696", 1, 200)
	if found != "" {
		t.Log("Found: ", found)
	} else {
		t.Error("Known dictionary hash not found: 36478381552696")
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
