package main

import (
	"fmt"
)

func main() {
	// create empty map
	// var colors map[string]string
	// colors := make(map[string]string)

	// [key is type string] value is type string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	fmt.Println(colors)

	printMap(colors)
}

// iterate through a map
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
