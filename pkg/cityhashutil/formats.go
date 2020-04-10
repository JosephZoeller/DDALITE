package cityhashutil

// HashCollision describes a hash - collision pair

type HashInParamsOffline struct {
	InputHashes []string
	HashType    string
	CompareFunc func([]string, func(string, string) bool)
}

type ColliderSpecifications struct {
	InputHashes []uint64 `json:"Hashes"`
	Dictionary  []string `json:"Dictionary"`
	Delimiter   string   `json:"Delimeter"`
	Depth       int      `json:"Depth"`
}

type ColliderResponse struct {
	Hashed   uint64 `json:"Hash"`
	Unhashed string `json:"Unhash"`
	Err      string `json:"errorString"`
}

type MessageResponse struct {
	Message string `json:"Message"`
}

type ClientSpecifications struct {
	InputHashes  []uint64   `json:"Hashes"`
	Dictionaries [][]string `json:"Dictionaries"`
	Delimiter    string     `json:"Delimeter"`
	Depth        int        `json:"Depth"`
}
