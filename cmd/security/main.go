// build: -darwin
package main

import (
	"fmt"
	"os"
)

type CommandFunc func([]string)

var commands map[string]CommandFunc

func registerCommand(commandName string, command CommandFunc) {
	if commands == nil {
		commands = make(map[string]CommandFunc)
	}
	commands[commandName] = command
}

func help() {
	fmt.Printf("%s: Emulates a subset of macOS' security command\n", os.Args[0])

	for subcommand, commandFunc := range commands {
		fmt.Printf("%s %s:\n", os.Args[0], subcommand)
		commandFunc([]string{"-help"})
	}
}

func main() {
	if len(os.Args) > 1 {
		commandName := os.Args[1]
		args := os.Args[2:]

		cmdFunc := commands[commandName]

		if cmdFunc != nil {
			cmdFunc(args)
		} else {
			help()
		}
	} else {
		help()
	}
}
