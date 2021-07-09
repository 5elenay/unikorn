package main

import (
	"fmt"
	"os"
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
	var repo Github

	if len(params) == 2 {
		// Username & Repo

		repo = Github{
			params[0],
			params[1],
			"main",
		}
	} else if len(params) >= 3 {
		// Username, Repo & Branch

		repo = Github{
			params[0],
			params[1],
			params[2],
		}
	} else {
		// Error

		OtherError("Please pass parameters correctly.")
	}

	DownloadFromGithub(repo)
}

// Remove Command
func CommandRemove(params []string) {
	if len(params) > 0 {
		// Remove Package from unikorn folder

		pkg := params[0]
		pkg = fmt.Sprintf("unikorn/%s", pkg)
		fmt.Printf("Trying to Remove Package From: %s\n", pkg)

		// Remove from Folder
		err := os.RemoveAll(pkg)
		UnexceptedError(err)

		fmt.Println("Removed the Package Successfully!")
	} else {
		// Error

		OtherError("Please pass a package name.")
	}
}
