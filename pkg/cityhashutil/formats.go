package cityhashutil

// HashCollision describes a hash - collision pair
type HashCollision struct {
	InputHash string `json:"hash"`
	Collision string `json:"collision"`
	Err       string `json:"errorStr"`
}

type HashInParams struct {
	InputHashes []string                                `json:"hashes"`
	HashType    string                                  `json:"HashType"` // should be an enum or something
	CompareFunc func([]string, func(string) bool) error `json:"compareFunction"`
}

type HashOutParams struct {
	Hashed   string `json:"hashedString"`
	Unhashed string `json:"unhashedString"`
	Err      string `json:"errorString"`
}
