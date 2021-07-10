package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Help Command
func CommandHelp(params []string) {
	if len(params) == 0 {
		// List All Commands
		fmt.Println("Please don't forget to check documentation on GitHub!\nList off all commands:")
		for _, item := range allCommands {
			fmt.Printf("    %s: %s\n", item.Name, item.Description)
		}
		fmt.Println("\nFor more information, Please type unikorn help <command>.")
		fmt.Printf("unikorn-%s | alpha-test\n", currentVersion)
	} else {
		// Find and Give Informations About Command
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
	if len(params) == 0 {
		// Error
		OtherError("Please pass a package name.")
	}

	pkg := params[0]
	pkg = fmt.Sprintf("unikorn/%s", pkg)
	fmt.Printf("Trying to Remove Package From: %s\n", pkg)

	// Remove from Folder
	err := os.RemoveAll(pkg)
	UnexceptedError(err)

	fmt.Println("Removed the Package Successfully!")
}

// Sync Command
func CommandSync(params []string) {
	if len(params) == 0 {
		// Error
		OtherError("Please pass a package name.")
	}

	// Save the Metadata Before Deleting
	pkg := params[0]
	metadata := fmt.Sprintf("unikorn/%s/unikorn.json", pkg)
	bytes, err := os.ReadFile(metadata)
	UnexceptedError(err)

	var saved PackageMetadata

	err = json.Unmarshal(bytes, &saved)
	UnexceptedError(err)

	// Run Remove Command
	CommandRemove(params)

	// Install Again
	CommandAdd(saved.Github)
}

// Find Command
func CommandFind(params []string) {
	if len(params) == 0 {
		// Error
		OtherError("Please pass a package name.")
	}

	// Get all Files & Folders in Directory
	files, err := os.ReadDir("unikorn")
	UnexceptedError(err)

	var metadatas []PackageMetadata

	for _, file := range files {
		if file.IsDir() {
			// Convert Metadata
			metadataFile := fmt.Sprintf("unikorn/%s/unikorn.json", file.Name())
			metadata := ConvertMetadata(metadataFile)

			// Append to Slice
			metadatas = append(metadatas, metadata)
		}
	}

	// Find And List Packages
	FindPackage(metadatas, params[0], func(found PackageMetadata, count int) {
		fmt.Printf("Result #%d\n    Package Name: %s\n    Package Description: %s\n    Package Tags: %s\n    PyPi Packages: %v\n\n", count, found.Name, found.Description, found.Tags, found.Pipreq)
	})
}

// Check Unikorn Update Command
func CommandUpdateCheck(params []string) {
	fmt.Printf("Checking For Updates... [Current Version: %s]\n", currentVersion)

	// Send Request and Get the Metadata
	response, err := http.Get("https://raw.githubusercontent.com/5elenay/unikorn/main/meta.json")
	UnexceptedError(err)

	defer response.Body.Close()

	var result UnikornMeta

	// Convert JSON to UnikornMeta Struct
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)
	UnexceptedError(err)

	// Check Version
	if result.Latest != currentVersion {
		fmt.Printf("Looks like you have an update for Unikorn. Please check: https://github.com/5elenay/unikorn/releases/latest\nLatest Release: %s\nCurrent: %s\n", result.Latest, currentVersion)
	} else {
		fmt.Println("Looks like you are using the latest version of Unikorn!")
	}
}
