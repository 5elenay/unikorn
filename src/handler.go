package main

import (
	"os"
	"strings"
)

var allCommands []Command

// Create All Commands
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
		{
			"remove",
			"Remove a package from project.",
			"unikorn remove <package name>",
			CommandRemove,
		},
		{
			"sync",
			"Sync a package (delete & install latest version.).",
			"unikorn sync <package name>",
			CommandSync,
		},
		{
			"find",
			"Find downloaded package(s) from name or tag.",
			"unikorn find <package name> | unikorn find <tag>",
			CommandFind,
		},
		{
			"check",
			"Check avaible updates for Unikorn.",
			"unikorn check",
			CommandUpdateCheck,
		},
	}
}

// Handle Commands
func HandleCommands() {
	parameters := os.Args[1:]

	// Check Parameter Length
	if len(parameters) == 0 {
		OtherError("You need to pass a parameter. Type 'help' for more information.")
	}

	parameter := strings.ToLower(parameters[0])

	// Find and Handle Command
	FindCommand(allCommands, parameter, func(found Command) {
		found.Handler(parameters[1:])
	})
}
