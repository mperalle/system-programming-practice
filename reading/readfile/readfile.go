package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify a file!")
		return
	}
	// Read file all at one and return contents in a []byte
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Read all file at once:\n", string(file))

	// Open file to read in chunk using buffer in byte slice
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Always think to close!
	defer f.Close()

	buffer := make([]byte, 16)

	for n := 0; err == nil; {
		n, err = f.Read(buffer)
		if err == nil {
			fmt.Print("Read in chunk: ", string(buffer[:n]))
		}
	}
	// Read returns EOF when no more content is available
	if err == io.EOF {
		fmt.Print(err, " returned, no more content available")
	} else {
		fmt.Println("\n\nError:", err)
	}

}
