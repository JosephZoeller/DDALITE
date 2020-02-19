package flyerdbutil

import (
	"testing"
)

func TestDB(t *testing.T) { // Apple == 4e1ec7e1511b (85894109417755)
	_, er := ConnectHard() // ASSUME: database is running on localhost:5432
	if er != nil {
		t.Error("Connection Failed - ", er)
	}
}