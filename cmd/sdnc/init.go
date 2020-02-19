package main

// Define ports for the different components that SDNC will be communicating with.
const (
	ColliderPort = "8080" // This is the expected port that the colliders will be listening on.
)

var (
	dictionaryLength int64 = 466550 //!! WARNING THIS IS ONLY TEMPORARY PLEASE ADD INIT LOGIC TO GET ACTUAL LENGTH FROM DB
	workerAddrs      []string
)

func init() {
	workerAddrs = make([]string, 0)
	/* What is this for?
	for i := 0; i < 2; i++ {
		wrkAddr := os.Getenv(fmt.Sprintf("wrkAddr_%d", 1))
		workerAddrs = append(workerAddrs, wrkAddr)
	}
	*/
}
