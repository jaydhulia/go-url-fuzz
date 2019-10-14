package requests

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-url-fuzz/config"
)

var client *http.Client = &http.Client{}

// Execute will be the main function that handles all the actual execution of requests and outputs
func Execute() error {
	file, err := os.Open(config.AppConfiguration.Wordlist)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var req *http.Request = &http.Request{Method: config.AppConfiguration.RequestType}
	for scanner.Scan() {
		err = doRequest(scanner.Text(), req)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

func doRequest(fuzzWord string, req *http.Request) error {
	rawURL := strings.Replace(config.AppConfiguration.URL, "FUZZ", fuzzWord, -1)
	URL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	req.URL = URL
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	go outputResult(res.StatusCode, rawURL)
	res.Body.Close()
	return nil
}
