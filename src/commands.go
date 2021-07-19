package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Help Command
func CommandHelp(params []string, _ []string) {
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

	GetConfirmation(fmt.Sprintf("Are you sure do you want to add this package (%s/%s ~ %s)?", repo.Username, repo.Repo, repo.Branch), options)
	DownloadFromGithub(repo, options)
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
		pkg := params[0]

		GetConfirmation(fmt.Sprintf("Are you sure do you want to delete this package (%s)?", pkg), options)

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
		GetConfirmation("Are you sure do you want to sync all packages?", options)

		// Get all Packages in the Unikorn Folder
		files, err := os.ReadDir("unikorn")
		UnexceptedError(err)

		for _, file := range files {
			if file.IsDir() {
				// Sync Package
				CommandSync([]string{file.Name()}, options)
			}
		}
	} else {
		// Save the Metadata Before Deleting
		pkg := params[0]
		metadata := fmt.Sprintf("unikorn/%s/unikorn.json", pkg)
		bytes, err := os.ReadFile(metadata)
		UnexceptedError(err)

		var saved PackageMetadata

		err = json.Unmarshal(bytes, &saved)
		UnexceptedError(err)

		GetConfirmation(fmt.Sprintf("Are you sure do you want to sync this package (%s)?", pkg), options)

		// Run Remove Command
		CommandRemove(params, options)

		// Install Again
		CommandAdd(saved.Github, options)
	}
}

// Find Command
func CommandFind(params []string, options []string) {
	if len(params) == 0 {
		// Error
		OtherError("Please pass a package name.")
	}

	// Get all Files & Folders in Directory
	metadatas := GetPackages()

	shouldFindAll := StringSliceContains(options, "all")

	// Find And List Packages
	FindPackage(metadatas, params[0], func(found PackageMetadata, count int) bool {
		fmt.Printf("Result #%d:\n    Package Name: %s\n    Package Description: %s\n    Package Tags: %s\n    PyPi Packages: %v\n\n", count, found.Name, found.Description, found.Tags, found.Pipreq)
		return shouldFindAll
	})
}

// Check Unikorn Update Command
func CommandUpdateCheck(_ []string, _ []string) {
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

// List All Packages Command
func CommandList(_ []string, _ []string) {
	metadatas := GetPackages()

	for index, item := range metadatas {
		fmt.Printf("Package #%d:\n    Package Name: %s\n    Package Description: %s\n    Package Tags: %s\n    PyPi Packages: %v\n\n", index+1, item.Name, item.Description, item.Tags, item.Pipreq)
	}
}

// Check Version Command
func CommandVersion(_ []string, _ []string) {
	fmt.Printf("unikorn-%s | alpha-test\n", currentVersion)
}

// Build a Template for Start Making Your Own Unikorn Package
func CommandNew(_ []string, options []string) {
	GetConfirmation("Are you sure you want to build new template? it may overwrite your unikorn.json and __init__.py file if already exists.", options)

	var packageName, packageDescription, packageTags, pypiRequirements, unikornPackages string

	// List of Question with Variable Pointers
	questions := []Question{
		{
			"Name of the Package: ",
			&packageName,
		},
		{
			"Description for your Package: ",
			&packageDescription,
		},
		{
			"Tags for your Package\n    (seperate with ',' character.): ",
			&packageTags,
		},
		{
			"If your Package uses some packages from pypi, List all of them\n    (seperate with ',' character.): ",
			&pypiRequirements,
		},
		{
			"If your Package uses some unikorn packages, List all of them\n    (seperate with spaces.)\n    (format is 'username:repo:branch' or 'username:repo'.): ",
			&unikornPackages,
		},
	}

	// Ask Questions
	for _, que := range questions {
		ReadInput(que.Question, que.Variable)
	}

	metadata := PackageMetadata{
		Name:        packageName,
		Description: packageDescription,
		Tags:        SeperateString(packageTags, ","),
		Pipreq:      SeperateString(pypiRequirements, ","),
		Github:      []string{},
	}

	// Check Tags and Pipreq
	if len(metadata.Tags) == 1 && metadata.Tags[0] == "" {
		metadata.Tags = []string{}
	}
	if len(metadata.Pipreq) == 1 && metadata.Pipreq[0] == "" {
		metadata.Pipreq = []string{}
	}

	// Format for Unireq
	var unikornList [][]string
	packages := strings.Split(unikornPackages, " ")

	for _, pkg := range packages {
		if pkg == "" {
			continue
		}

		unikornList = append(unikornList, strings.Split(pkg, ":"))
	}

	if len(unikornList) == 0 {
		metadata.Unireq = [][]string{}
	} else {
		metadata.Unireq = unikornList
	}

	// Convert Metadata to Beautify JSON
	data, err := json.MarshalIndent(metadata, "", "    ")
	UnexceptedError(err)

	// Create unikorn.json
	fmt.Println("Creating unikorn.json File...")
	err = os.WriteFile("unikorn.json", data, 0666)
	UnexceptedError(err)
	fmt.Println("Created unikorn.json File Successfully!")

	// Create Package Folder
	fmt.Println("Creating src Folder...")
	os.Mkdir("src", os.ModePerm)
	fmt.Println("Created src Folder Successfully!")

	// Add __init__.py File
	fmt.Println("Creating __init__.py File...")
	err = os.WriteFile("./src/__init__.py", []byte("hello = lambda x: print('Hello: ', x)"), 0666)
	UnexceptedError(err)
	fmt.Println("Created __init__.py File Successfully!")
}
