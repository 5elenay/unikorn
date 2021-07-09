package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var currentVersion string = "0.0.1"

func main() {
	// Handle Commands
	HandleCommands()
}

// Download File from URL
func DownloadFile(url string, name string, temp string) string {
	resp, err := http.Get(url)
	UnexceptedError(err)

	// Close When All the Jobs are Finished
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	UnexceptedError(err)

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", temp, name), bodyBytes, 0666)
	UnexceptedError(err)

	return name
}

// Download a Github Repo
func DownloadFromGithub(github Github) {
	// Download Zip URL
	repoUrl := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", github.Username, github.Repo, github.Branch)

	// File Name
	fileName := fmt.Sprintf("%s.zip", github.Repo)

	// Create Temp Directory for ZIP
	folderName := CreateTempDirectory()

	// Download the ZIP from URL
	DownloadFile(repoUrl, fileName, folderName)
}

func CreateTempDirectory() string {
	// Create Temp Directory for Unikorn
	tempFolder := ".unik"

	err := os.Mkdir(tempFolder, os.ModePerm)
	UnexceptedError(err)

	return tempFolder
}

// Find Command from List of Commands
func FindCommand(commands []Command, name string, function func(found Command)) {
	var found bool

	for _, item := range commands {
		// Check if has Same Command Name
		if item.Name == name {
			// Run the Function
			function(item)

			// Make Found True and Break the Loop
			found = true
			break
		}
	}

	// Check if Found or Not
	CommandNotFound(found)
}
