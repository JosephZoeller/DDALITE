package cityhashutil

// HashCollision describes a hash - collision pair

type HashInParamsOffline struct {
	InputHashes []string
	HashType    string
	CompareFunc func([]string, func(string, string) bool)
}

type ColliderSpecifications struct {
	InputHashes []string `json:"Hashes"`
	Dictionary  []string `json:"Dictionary"`
	Delimiter   string   `json:"Delimeter"`
	Depth       int      `json:"Depth"`
}

type ColliderResponse struct {
	Hashed   string `json:"hashedString"`
	Unhashed string `json:"unhashedString"`
	Err      string `json:"errorString"`
}

type MessageResponse struct {
	Message string `json:"Message"`
}

type ClientSpecifications struct {
	InputHashes  []string   `json:"Hashes"`
	Dictionaries [][]string `json:"Dictionaries"`
	Delimiter    string     `json:"Delimeter"`
	Depth        int        `json:"Depth"`
}
