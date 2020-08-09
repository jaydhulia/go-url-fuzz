package requests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-url-fuzz/config"
	"github.com/valyala/fasthttp"
)

var (
	// Dummy channel to control number of concurrent requests
	concurrentRequests chan struct{}
	// The waitgroup - to ensure all go routines complete before main exits
	wg *sync.WaitGroup = &sync.WaitGroup{}
)

// Execute will be the main function that handles all the actual execution of requests and outputs
func Execute() error {
	// set the number of concurrentRequests channel to the concurrency requested
	concurrentRequests = make(chan struct{}, config.AppConfiguration.Concurrency)

	file, err := os.Open(config.AppConfiguration.Wordlist)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wg.Add(1)
		go doRequest(scanner.Text())
	}
	// Wait for all the requests and their outputs to complete
	wg.Wait()
	return nil
}

func doRequest(fuzzWord string) error {

	defer wg.Done()

	rawURL := strings.Replace(config.AppConfiguration.URL, "FUZZ", fuzzWord, -1)

	// Aquire the request and response from the pool
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	// Set the required static headers and URL
	req.Header.SetMethod(config.AppConfiguration.RequestType)
	req.SetRequestURI(rawURL)

	// Wait for the request's turn using chan to ensure max concurrent connections
	concurrentRequests <- struct{}{}
	err := fasthttp.Do(req, resp)
	<-concurrentRequests

	fasthttp.ReleaseRequest(req)

	if err != nil {
		return err
	}

	wg.Add(1)
	go outputResult(resp.StatusCode(), rawURL)

	fasthttp.ReleaseResponse(resp)
	return nil
}
