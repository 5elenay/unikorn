package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Download File from URL
func DownloadFile(url, name, temp string) string {
	resp, err := http.Get(url)
	UnexceptedError(err)

	// Close When All the Jobs are Finished
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	UnexceptedError(err)

	filePath := fmt.Sprintf("%s/%s", temp, name)

	err = ioutil.WriteFile(filePath, bodyBytes, 0666)
	UnexceptedError(err)

	return filePath
}

// Download a Github Repo
func DownloadFromGithub(github Github) {
	// Download Zip URL
	repoUrl := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", github.Username, github.Repo, github.Branch)
	fmt.Println("Repo URL:", repoUrl)

	// File Name
	fileName := fmt.Sprintf("%s.zip", github.Repo)

	// Create Temp Directory for ZIP
	fmt.Println("Creating .unik Folder.")
	folderName := CreateTempDirectory()
	fmt.Println("Created .unik Folder.")

	// Download the ZIP from URL
	fmt.Println("Downloading the File:", repoUrl)
	filePath := DownloadFile(repoUrl, fileName, folderName)
	fmt.Println("Downloaded With Name:", fileName)

	// Extract the ZIP
	fmt.Println("Extracting the ZIP...")
	ExtractZip(filePath, folderName)
	fmt.Println("Extract Finished!")

	// Delete the ZIP
	fmt.Println("Removing the ZIP.")
	err := os.Remove(filePath)
	UnexceptedError(err)

	fmt.Println("Removed Successfully!")

	// Get Extracted Folder Name
	extractedName := fmt.Sprintf("%s-%s", github.Repo, github.Branch)

	// Create Unikorn Folder
	fmt.Println("Creating Unikorn Folder.")
	unikornPath := CreateUnikornDirectory()
	fmt.Println("Created Unikorn Folder.")

	// Convert unikorn.json to PackageMetadata
	fmt.Println("Reading the unikorn.json")
	convertedData := ConvertMetadata(fmt.Sprintf(".unik/%s/unikorn.json", extractedName))
	fmt.Println("Converted Metadata Successfully")

	// Move Folder with Metadata
	RenameAndMove(github, convertedData, unikornPath, extractedName)

	// Delete .unix Folder
	fmt.Println("Deleting the .unik Folder..")
	os.Remove(folderName)
	fmt.Println("Deleted .unik Folder Successfully!")

	// Check Pypi Packages
	fmt.Println("Checking for PyPi Packages...")
	convertedData.CheckPackages()
}

func RenameAndMove(github Github, metadata PackageMetadata, path, old string) {
	folderName := metadata.Name

	if metadata.Name == "" {
		folderName = github.Repo
	}

	oldPath := fmt.Sprintf(".unik/%s/src", old)
	newPath := fmt.Sprintf("%s/%s", path, folderName)

	// Move Files With Rename
	fmt.Printf("Moving Folder %s to the %s...\n", oldPath, newPath)
	os.Rename(oldPath, newPath)

	fmt.Println("Moved Successfully!")
}

func CreateTempDirectory() string {
	// Create Temp Directory for Unikorn
	tempFolder := ".unik"
	os.Mkdir(tempFolder, os.ModePerm)

	return tempFolder
}

func CreateUnikornDirectory() string {
	// Create Unikorn Directory for Packages
	folder := "unikorn"
	os.Mkdir(folder, os.ModePerm)

	return folder
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

// Extract Zip
func ExtractZip(path, dest string) {
	// Open New ZIP Reader
	reader, err := zip.OpenReader(path)
	UnexceptedError(err)

	defer reader.Close()

	// Read all Files in the ZIP
	for _, file := range reader.File {
		// Clone File
		CloneFile(file, dest)
	}
}

// Converter for unikorn.json
func ConvertMetadata(path string) PackageMetadata {
	// Read the File
	fileData, err := ioutil.ReadFile(path)
	UnexceptedError(err)

	// Decode JSON
	pkg := PackageMetadata{}
	err = json.Unmarshal(fileData, &pkg)
	UnexceptedError(err)

	return pkg
}

// Clone File
func CloneFile(file *zip.File, dest string) {
	// Create Full Path
	path := filepath.Join(dest, file.Name)
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	UnexceptedError(err)

	// Clone the File
	reader, err := file.Open()
	UnexceptedError(err)

	defer reader.Close()

	// Check if File
	if !file.FileInfo().IsDir() {
		// Copy the File
		newFile, err := os.Create(path)

		UnexceptedError(err)
		_, err = io.Copy(newFile, reader)
		UnexceptedError(err)

		newFile.Close()
	}
}

func GetFileName(file *zip.File) string {
	fullPath := file.Name
	splittedPath := strings.Split(fullPath, "/")

	var filtederPath []string

	for _, item := range splittedPath {
		if item != "" {
			filtederPath = append(filtederPath, item)
		}
	}

	fileName := filtederPath[len(filtederPath)-1]

	return fileName
}
