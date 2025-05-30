// build: -darwin
package main

import (
	"log"
	"os"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Error initializing clipboard: %v", err)
		os.Exit(1)
	}

	os.Stdout.Write(clipboard.Read(clipboard.FmtText))
}
