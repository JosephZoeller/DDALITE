package cityhashutil

// HashCollision describes a hash - collision pair
type HashCollision struct {
	InputHash string `json:"hash"`
	Collision string `json:"collision"`
	Err       string `json:"errorStr"`
}

type HashInParamsOffline struct {
	InputHashes []string
	HashType    string
	CompareFunc func([]string, func(string, string) bool)
}

type HashInParamsOnline struct {
	InputHashes []string `json:"Hashes"`
	Dictionary  []string `json:"Dictionary"`
	Delimiter   string   `json:"Delimeter"`
	Depth       int      `json:"Depth"`
}

type HashOutParams struct {
	Hashed   string `json:"hashedString"`
	Unhashed string `json:"unhashedString"`
	Err      string `json:"errorString"`
}

type ResponseMessage struct {
	Message string `json:"Message"`
}

type ClientPost struct {
	InputHashes  []string   `json:"Hashes"`
	Dictionaries [][]string `json:"Dictionaries"`
	Delimiter    string     `json:"Delimeter"`
	Depth        int        `json:"Depth"`
}
