package main // all files in this project must belong in a package
// main = executable type package -> "go build main.go" will output an executable file (main.exe)
// any other name = reusable type package -> "go build <name>.go" will output code dependancies / libaries / helper functions

import "fmt"

// executable package must have a main function
func main() {
	fmt.Println("Hello World")
}
