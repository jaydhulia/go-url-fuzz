package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Configuration describes the configuration of fuzzer
type Configuration struct {
	RequestType string
	URL         string
	Wordlist    string
	Concurrency int
}

var (
	defaultRequestType       string = "GET"
	defaultConcurrentThreads int    = 100
	// AppConfiguration defines the app configuration
	AppConfiguration Configuration = Configuration{}
)

// ParseFlags parses the CLI flags and defines the application configuration
func ParseFlags() error {
	flag.StringVar(&AppConfiguration.RequestType, "X", defaultRequestType, "-X=\"GET\"")
	flag.StringVar(&AppConfiguration.URL, "url", "", "-url=\"https://example.com/FUZZ\"")
	flag.StringVar(&AppConfiguration.Wordlist, "wordlist", "", "-wordlist=PATH/TO/FILE")
	flag.IntVar(&AppConfiguration.Concurrency, "c", defaultConcurrentThreads, "-c=100")
	flag.Parse()
	if AppConfiguration.RequestType != "GET" {
		return fmt.Errorf("sorry, currently only GET requests are supported")
	}
	if AppConfiguration.URL == "" || !strings.Contains(AppConfiguration.URL, "FUZZ") {
		return fmt.Errorf("url with 'FUZZ' is required. Please see usaage")
	}
	if AppConfiguration.Concurrency < 1 || AppConfiguration.Concurrency > 200 {
		return fmt.Errorf("please define a reasonable concurrency")
	}

	if _, err := os.Stat(AppConfiguration.Wordlist); err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
