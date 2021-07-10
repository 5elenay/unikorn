package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Help Command
func CommandHelp(params []string, options []string) {
	if len(params) == 0 {
		fmt.Println("Please don't forget to check documentation on GitHub!\nList off all commands:")
		// List All Commands
		for _, item := range allCommands {
			fmt.Printf("    %s: %s\n", item.Name, item.Description)
		}
		fmt.Println("\nFor more information, Please type 'unikorn help <command>'.")
		fmt.Printf("unikorn-%s | alpha-test\n", currentVersion)
	} else {
		// Find and Give Informations About Command
		param := strings.ToLower(params[0])

		FindCommand(allCommands, param, func(found Command) {
			fmt.Printf("Description:\n    %s\nUsage:\n    %s\n", found.Description, strings.Join(found.Usage, "\n    "))

			// List all Options
			if len(found.Options) > 0 {
				fmt.Print("Options:")

				for _, option := range found.Options {
					fmt.Printf("\n    %s:\n        %s\n", option.Name, option.Description)
				}
			}
		})
	}
}

// Download Command
func CommandAdd(params []string, options []string) {
	var repo Github

	if len(params) == 0 {
		// Read Unipkg File
		bytes, err := os.ReadFile("unipkg")
		UnexceptedError(err)

		data := string(bytes)
		splitted := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")

		// Split Lines
		for _, item := range splitted {
			// Split Items
			args := strings.Split(item, " ")

			// Run Command
			CommandAdd(args, options)
		}

		return
	} else if len(params) == 2 {
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

	GetConfirmation("Are you sure do you want to add this package?", options)
	DownloadFromGithub(repo)
}

// Remove Command
func CommandRemove(params []string, options []string) {
	if len(params) == 0 {
		GetConfirmation("Are you sure do you want to delete all the packages?", options)

		fmt.Println("Trying to remove all of the packages")

		// Remove from Folder
		err := os.RemoveAll("unikorn")
		UnexceptedError(err)

		fmt.Println("Removed the Packages Successfully!")
	} else {
		GetConfirmation("Are you sure do you want to delete this package?", options)

		pkg := params[0]
		pkg = fmt.Sprintf("unikorn/%s", pkg)
		fmt.Printf("Trying to Remove Package From: %s\n", pkg)

		// Remove from Folder
		err := os.RemoveAll(pkg)
		UnexceptedError(err)

		fmt.Println("Removed the Package Successfully!")
	}
}

// Sync Command
func CommandSync(params []string, options []string) {
	if len(params) == 0 {
		// Error
		OtherError("Please pass a package name.")
	}

	GetConfirmation("Are you sure do you want to sync this package?", options)

	// Save the Metadata Before Deleting
	pkg := params[0]
	metadata := fmt.Sprintf("unikorn/%s/unikorn.json", pkg)
	bytes, err := os.ReadFile(metadata)
	UnexceptedError(err)

	var saved PackageMetadata

	err = json.Unmarshal(bytes, &saved)
	UnexceptedError(err)

	// Run Remove Command
	CommandRemove(params, options)

	// Install Again
	CommandAdd(saved.Github, options)
}

// Find Command
func CommandFind(params []string, options []string) {
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
func CommandUpdateCheck(_ []string, options []string) {
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

// Initialize
func CommandInit(_ []string, options []string) {
	GetConfirmation("Are you sure you want to initialize? it may delete your unipkg file if already exists.", options)

	// Create Unikorn Directory
	CreateUnikornDirectory()

	// Create Unipkg File
	file, err := os.Create("unipkg")
	UnexceptedError(err)

	file.Close()
}
