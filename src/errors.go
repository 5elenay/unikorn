// Fatals
/*
	1 - Unexcepted Error
	2 - Command not Found
	3 - Unknown (Other)
*/

package main

import (
	"log"
	"os"
)

// Raises When Unexcepted Error Found
func UnexceptedError(err error) {
	if err != nil {
		log.Fatal("Unexcepted Error: ", err)
		os.Exit(1)
	}
}

// Raises When Command not Found
func CommandNotFound(found bool) {
	if !found {
		log.Fatal("Command not Found!")
		os.Exit(2)
	}
}

// Other Errors (Custom Errors)
func OtherError(desc string) {
	log.Fatal(desc)
	os.Exit(3)
}
