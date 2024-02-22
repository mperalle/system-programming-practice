package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Read file all at one and return contents in a []byte
	file, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(file))
}
