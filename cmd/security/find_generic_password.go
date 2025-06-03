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
		)

		// Parse flags
		findPasswordFlags.StringVar(&a, "a", "", "account to match")
		findPasswordFlags.StringVar(&s, "s", "", "service to match string")

		findPasswordFlags.Parse(args)

		if s == "" {
			log.Fatal("Empty service name")
		}

		if a == "" {
			log.Fatal("Empty account name")
		}

		password, e := keyring.Get(s, a)
		if e != nil {
			log.Fatalf("Can't find password for %s - %s: %v", s, a, e)
		}

		fmt.Print(password)
	}
	registerCommand(funcName, findGenericPassword)
}
