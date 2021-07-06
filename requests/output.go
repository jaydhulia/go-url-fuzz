package requests

import (
	"fmt"

	"github.com/go-url-fuzz/config"
	"github.com/gookit/color"
)

var (
	// Dummy channel to ensure output doesn't get messed up with concurrent requests
	outputChan chan struct{}
)

func init() {

	outputChan = make(chan struct{}, 1)
}

// OutputResult outputs the result of each request
// TODO: add more colors based on responses
func outputResult(responseCode int, URL string) {
	defer wg.Done()
	// if user has specified a show codes flag
	if config.AppConfiguration.ShowCodesMap != nil {
		// code is not in allowed list
		if _, ok := config.AppConfiguration.ShowCodesMap[responseCode]; !ok {
			return
		}
	}
	red := color.FgRed.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render

	// Wait for the output's turn using chan to ensure output doesn't get messed up
	outputChan <- struct{}{}
	fmt.Printf("[ Status: ")
	if responseCode >= 200 && responseCode < 300 {
		fmt.Printf("%s", green(responseCode))
	} else {
		fmt.Printf("%s", red(responseCode))
	}
	fmt.Printf(" \t Request: %-50s ]\n", cyan(URL))
	<-outputChan
}
