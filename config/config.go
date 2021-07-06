package config

// Configuration describes the configuration of fuzzer
type Configuration struct {
	RequestType  string
	URL          string
	Wordlist     string
	Concurrency  int
	ShowCodes    []int
	ShowCodesMap map[int]int
}

var (
	DefaultRequestType       string = "GET"
	DefaultConcurrentThreads int    = 100
	DefaultShowCodes         []int  = make([]int, 0)
	// AppConfiguration defines the app configuration
	AppConfiguration Configuration = Configuration{}
)
