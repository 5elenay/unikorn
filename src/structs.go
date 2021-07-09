package main

import (
	"fmt"
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
	Name, Description string
	Tags, Pipreq      []string
}

func (metadata PackageMetadata) CheckPackages() {
	if len(metadata.Pipreq) > 0 {
		packages := fmt.Sprintf("pip install %s", strings.Join(metadata.Pipreq, " "))
		fmt.Printf("\nThis package requires some packages from PyPi.\nFor download all of them:\n    %s", packages)
	} else {
		fmt.Println("Looks like this package does not contains any PyPi package!")
	}
}
