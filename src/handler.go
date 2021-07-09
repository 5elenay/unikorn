package main

import (
	"os"
	"strings"
)

var allCommands []Command

// Create Commands
func init() {
	allCommands = []Command{
		{
			"help",
			"Shows list of all commands.",
			"unikorn help | unikorn help <command>",
			CommandHelp,
		},
		{
			"add",
			"Download & add a package from github.",
			"unikorn add <github username> <repo name> | unikorn add <github username> <repo name> <branch>",
			CommandAdd,
		},
	}
}

// Handle Commands
func HandleCommands() {
	parameters := os.Args[1:]

	if len(parameters) == 0 {
		OtherError("You need to pass a parameter. Type 'help' for more information.")
	}

	parameter := strings.ToLower(parameters[0])

	FindCommand(allCommands, parameter, func(found Command) {
		found.Handler(parameters[1:])
	})
}
