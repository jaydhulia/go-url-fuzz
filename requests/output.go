package requests

import (
	"fmt"

	"github.com/gookit/color"
)

// OutputResult outputs the result of each request
// TODO: add more colors based on responses
func outputResult(responseCode int, URL string) {
	red := color.FgRed.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render
	fmt.Printf("[ Status: ")
	if responseCode == 200 || responseCode == 201 || responseCode == 202 {
		fmt.Printf("%s", green(responseCode))
	} else {
		fmt.Printf("%s", red(responseCode))
	}
	fmt.Printf(" \t Request: %-50s ]\n", cyan(URL))
}
