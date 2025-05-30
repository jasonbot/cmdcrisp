// build: -darwin
package main

import (
	"io"
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

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin: %v", err)
		os.Exit(1)
	}

	clipboard.Write(clipboard.FmtText, stdin)
}
