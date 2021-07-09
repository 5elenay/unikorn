package main

import (
	"fmt"
	"strings"
)

// Help Command
func CommandHelp(params []string) {
	if len(params) == 0 {
		fmt.Println("Please don't forget to check documentation on GitHub!\nList off all commands:")
		for _, item := range allCommands {
			fmt.Printf("    %s: %s\n", item.Name, item.Description)
		}
		fmt.Println("\nFor more information, Please type unikorn help <command>.")
		fmt.Printf("unikorn-%s | alpha-test\n", currentVersion)
	} else {
		param := strings.ToLower(params[0])

		FindCommand(allCommands, param, func(found Command) {
			fmt.Printf("Description: %s\nUsage: %s\n", found.Description, found.Usage)
		})
	}
}

// Download Command
func CommandAdd(params []string) {
	if len(params) == 2 {
		// Username & Repo

		fmt.Println("Soon!")
	} else if len(params) >= 3 {
		// Username, Repo & Branch

		fmt.Println("Soon!")
	} else {
		// Error

		OtherError("Please pass parameters correctly.")
	}
}
