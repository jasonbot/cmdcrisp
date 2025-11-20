package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zalando/go-keyring"
)

func init() {
	funcName := "find-generic-password"
	findGenericPassword := func(args []string) {
		findPasswordFlags := flag.NewFlagSet(funcName, flag.ExitOnError)

		// Define flags
		var (
			a string // account to match
			s string // service to match string
			w bool   // Display the password(only) for the item found
		)

		// Parse flags
		findPasswordFlags.StringVar(&a, "a", "", "account to match")
		findPasswordFlags.StringVar(&s, "s", "", "service to match string")
		findPasswordFlags.BoolVar(&w, "w", true, "display the password")

		findPasswordFlags.Parse(args)

		if s == "" {
			log.Fatal("Empty service name")
		}

		if a == "" {
			log.Fatal("Empty account name")
		}

		password, e := keyring.Get(s, a)
		if e != nil {
			KeyringGet(s, a)
		}

		if w {
			fmt.Print(password)
		}
	}
	registerCommand(funcName, findGenericPassword)
}
