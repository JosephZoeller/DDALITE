package main


import (
	"fmt"
	"testing"

	"github.com/200106-uta-go/JKJP2/pkg/flyerdbutil"
)

func Test(t *testing.T) { // TENSION_NECK == 60db807ad6c6 (106495869638342) // apple == b10f405db64b (194679062509131) Apple == 4e1ec7e1511b (85894109417755)
	found := findCollisionFile("194679062509131", 0, 450000)
	if found != "" {
		fmt.Println(found)
	} else {
		t.Error("Known dictionary hash not found: 194679062509131")
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

	found := findDBCollision("85894109417755", 0, 450000)
	if found != "" {
		fmt.Println(found)
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

		row := flyerDB.QueryRow("SELECT word FROM dic WHERE idx = $1", 0)
		er = row.Scan(&word)
		if er != nil {
			t.Error("Query Failed - ", er)
		}

}