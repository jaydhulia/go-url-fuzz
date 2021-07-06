package config

// Configuration describes the configuration of fuzzer
type Configuration struct {
	RequestType string
	URL         string
	Wordlist    string
	Concurrency int
}

var (
	DefaultRequestType       string = "GET"
	DefaultConcurrentThreads int    = 100
	// AppConfiguration defines the app configuration
	AppConfiguration Configuration = Configuration{}
)
