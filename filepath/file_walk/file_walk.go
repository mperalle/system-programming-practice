package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// Ensure path is specified
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path!")
		os.Exit(2)
	}
	// Get absolute path
	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}

	var c struct {
		files int
		dirs  int
	}
	fmt.Println("Listing files in", root)

	// Walk the tree to count files and folders
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})
	fmt.Printf("Total: %d files in %d directories\n", c.files, c.dirs)
}
