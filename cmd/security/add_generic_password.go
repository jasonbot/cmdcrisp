package main

import (
	"flag"
	"log"

	"github.com/zalando/go-keyring"
)

func init() {
	funcName := "add-generic-password"

	addGenericPassword := func(args []string) {
		addPasswordFlags := flag.NewFlagSet(funcName, flag.ExitOnError)

		// Define flags
		var (
			a string // account to match
			c string // creator to match (four-character code)
			D string // kind to match string
			s string // service to match string
			w string // Specify password to be added
		)

		// Parse flags
		addPasswordFlags.StringVar(&a, "a", "", "account to match")
		addPasswordFlags.StringVar(&c, "c", "", "creator to match (four-character code)")
		addPasswordFlags.StringVar(&D, "D", "", "kind to match string")
		addPasswordFlags.StringVar(&s, "s", "", "service to match string")
		addPasswordFlags.StringVar(&w, "w", "", "password to set")

		addPasswordFlags.Parse(args)

		if s == "" {
			log.Fatal("Empty service name")
		}

		if a == "" {
			log.Fatal("Empty account name")
		}

		keyring.Set(s, a, w)
	}
	registerCommand(funcName, addGenericPassword)
}
