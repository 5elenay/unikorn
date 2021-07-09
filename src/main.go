package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var currentVersion string = "0.0.1"

func main() {
	// Handle Commands
	HandleCommands()
}

// Download File from URL
func DownloadFile(url string, name string) {
	resp, err := http.Get(url)
	UnexceptedError(err)

	// Close When All the Jobs are Finished
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	UnexceptedError(err)

	ioutil.WriteFile(name, bodyBytes, 0666)
}

// Download a Github Repo
func DownloadFromGithub(github Github) {
	var branch string

	if github.Branch == "" {
		branch = "master"
	} else {
		branch = github.Branch
	}

	// Download Zip URL
	repoUrl := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", github.Username, github.Repo, branch)

	// File Name
	fileName := fmt.Sprintf("%s.zip", github.Repo)

	DownloadFile(repoUrl, fileName)
}

// Find Command
func FindCommand(commands []Command, name string, function func(found Command)) {
	var found bool

	for _, item := range commands {
		if item.Name == name {
			function(item)
			found = true
			break
		}
	}

	CommandNotFound(found)
}
