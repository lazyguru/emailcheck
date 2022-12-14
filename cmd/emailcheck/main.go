package main

import (
	"github.com/lazyguru/emailcheck/internal/emailcheck"
	"log"
)

func main() {
	err := emailcheck.Initialize()
	if err != nil {
		log.Fatalf("Unable to initialize client: %v", err)
	}

	emailcheck.Run()
}
