package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// Github Repo Struct
type Github struct {
	Username, Repo, Branch string
}

// Unikorn Command Struct
type Command struct {
	Name, Description, Usage string
	Handler                  func(params []string)
}

// Package Metadata Struct
type PackageMetadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Pipreq      []string `json:"pipreq"`
	Github      []string `json:"github"`
}

// Create Archive Url from Github Struct.
func (github Github) CreateUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", github.Username, github.Repo, github.Branch)
}

// Check Packages
func (metadata PackageMetadata) CheckPackages() {
	if len(metadata.Pipreq) > 0 {
		packages := fmt.Sprintf("pip install %s", strings.Join(metadata.Pipreq, " "))
		fmt.Printf("\nThis package requires some packages from PyPi.\nFor download all of them:\n    %s", packages)
	} else {
		fmt.Println("Looks like this package does not contains any PyPi package!")
	}
}

// Add Github Details
func (metadata PackageMetadata) UpdateGithubDetails(data []string, path string) {
	metadata.Github = data

	bytes, err := json.Marshal(metadata)
	if err != nil {
		UnexceptedError(err)
	}

	ioutil.WriteFile(path, bytes, 0666)
}
