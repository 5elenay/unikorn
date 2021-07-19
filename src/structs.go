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
	Name, Description string
	Usage             []string
	Options           []Option
	Handler           func(params []string, options []string)
}

// Unikorn Command Line Option Struct
type Option struct {
	Name, Description string
}

// Unikorn Metadata Struct
type UnikornMeta struct {
	Latest, License, Git string
}

// Package Metadata Struct
type PackageMetadata struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Pipreq      []string   `json:"pipreq"`
	Unireq      [][]string `json:"unireq"`
	Github      []string   `json:"github"`
}

// Question Struct
type Question struct {
	Question string
	Variable *string
}

// Create Archive Url from Github Struct.
func (github Github) CreateUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", github.Username, github.Repo, github.Branch)
}

// Check Packages
func (metadata PackageMetadata) CheckPackages(options []string) {
	// Check Unikorn Packages
	if len(metadata.Unireq) > 0 {
		for _, item := range metadata.Unireq {
			CommandAdd(item, options)
		}
	}

	// Check PyPi Packages
	if len(metadata.Pipreq) > 0 {
		packages := fmt.Sprintf("pip install %s", strings.Join(metadata.Pipreq, " "))
		fmt.Printf("\nThis (%s) package requires some packages from PyPi.\nFor download all of them:\n    %s\n\n", metadata.Name, packages)
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
