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
	Words       int      `json:"Depth"`
	StartsWith  string   `json:"StartsWith"`
	EndsWith	string	`json:"EndsWith"`
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
	Colliders []ColliderSpecifications `json:"Colliders"`
}
