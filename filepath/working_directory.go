package main

import (
	"fmt"
	"os"
)

// changing working directory:
// - simplifiy file path management
// - isolation of the program's part = avoid accidental file manipulations
// - access control: limit access to only authorized areas

func main() {
	// Return the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error returning current working directory:", err)
		return
	}
	fmt.Println("Starting directory:", wd)

	// Change current directory
	if err := os.Chdir("/Users/Guest"); err != nil {
		fmt.Println("Error changing current directory:", err)
		return
	}

	// Return the current working directory
	if wd, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Changing directory to:", wd)

	// Reading files in directory
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
	}

	for _, file := range files {
		fmt.Println("File in directory:", file.Name())
	}

}
