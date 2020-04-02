package cityhashutil

// HashCollision describes a hash - collision pair
type HashCollision struct {
	InputHash string `json:"hash"`
	Collision string `json:"collision"`
	Err       string `json:"errorStr"`
}

type HashInParams struct {
	InputHashes []string                            
	HashType    string                                   
	CompareFunc func([]string, func(string, string) bool)
}

type HashOutParams struct {
	Hashed   string `json:"hashedString"`
	Unhashed string `json:"unhashedString"`
	Err      string `json:"errorString"`
}