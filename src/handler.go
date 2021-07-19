package main

import (
	"os"
	"strings"
)

var allCommands []Command
var noConfirmationOption Option = Option{
	"no-confirmation",
	"Skip confirmation process for command.",
}

// Create All Commands
func init() {
	allCommands = []Command{
		{
			"help",
			"Shows list of all commands.",
			[]string{"unikorn help", "unikorn help <command>"},
			[]Option{},
			CommandHelp,
		},
		{
			"add",
			"Download & add a package from github.",
			[]string{"unikorn add", "unikorn add <github username> <repo name>", "unikorn add <github username> <repo name> <branch>"},
			[]Option{
				noConfirmationOption,
			},
			CommandAdd,
		},
		{
			"remove",
			"Remove a package from project.",
			[]string{"unikorn remove", "unikorn remove <package name>"},
			[]Option{
				noConfirmationOption,
			},
			CommandRemove,
		},
		{
			"sync",
			"Sync a package (delete & install latest version.).",
			[]string{"unikorn sync <package name>"},
			[]Option{
				noConfirmationOption,
			},
			CommandSync,
		},
		{
			"find",
			"Find downloaded package(s) from name or tag.",
			[]string{"unikorn find <package name>", "unikorn find <tag>"},
			[]Option{
				{
					"all",
					"Find all packages.",
				},
			},
			CommandFind,
		},
		{
			"list",
			"List downloaded packages.",
			[]string{"unikorn list"},
			[]Option{},
			CommandList,
		},
		{
			"check",
			"Check avaible updates for Unikorn.",
			[]string{"unikorn check"},
			[]Option{},
			CommandUpdateCheck,
		},
		{
			"init",
			"Initialize basic setup for Unikorn.",
			[]string{"unikorn init"},
			[]Option{
				noConfirmationOption,
			},
			CommandInit,
		},
		{
			"version",
			"Check Unikorn version.",
			[]string{"unikorn version"},
			[]Option{},
			CommandVersion,
		},
		{
			"new",
			"Generate a Unikorn package template in <10 second.",
			[]string{"unikorn new"},
			[]Option{},
			CommandNew,
		},
	}
}

// Handle Commands
func HandleCommands() {
	args := os.Args[1:]

	// Check Parameter Length
	if len(args) == 0 {
		OtherError("You need to pass a argument. Type 'help' for more information.")
	}

	var params []string
	var options []string

	for _, item := range args {
		lowerCaseItem := strings.ToLower(item)

		if strings.HasPrefix(lowerCaseItem, "-") {
			options = append(options, strings.TrimPrefix(lowerCaseItem, "-"))
		} else {
			params = append(params, lowerCaseItem)
		}
	}

	// Check Parameter Length
	if len(params) == 0 {
		OtherError("You need to pass a parameter. Type 'help' for more information.")
	}

	parameter := strings.ToLower(params[0])

	// Find and Handle Command
	FindCommand(allCommands, parameter, func(found Command) {
		found.Handler(params[1:], options)
	})
}
