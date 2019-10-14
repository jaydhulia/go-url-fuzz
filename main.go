package main

import (
	"log"

	"github.com/go-url-fuzz/config"
	"github.com/go-url-fuzz/requests"
)

func main() {
	if err := config.ParseFlags(); err != nil {
		log.Fatalf(err.Error())
	}
	if err := requests.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
