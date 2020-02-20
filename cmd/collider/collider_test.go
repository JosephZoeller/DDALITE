package main


import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	found := findCollisionFile("4132412743", 0, 450000)
	if found != "" {
		fmt.Println(found)
	} else {
		t.Error("Known dictionary hash not found: 41320793412743")
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findCollisionFile("1", 0, i)
	}

}