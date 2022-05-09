package main

import (
	"fmt"
	"io"
	"os"
)

// Run the command -> go run main.go myfile.txt
func main() {
	// Get the filename form std input
	fmt.Println(os.Args)
	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// os.Stdout implements the Writer interface
	// File implements the Reader interface
	io.Copy(os.Stdout, file)
}
