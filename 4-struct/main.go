package main

import "fmt"

// Define a new struct
type person struct {
	// <property name> <type>
	firstName   string
	lastName    string
	contactInfo // Embedding struct
}

type contactInfo struct {
	email   string
	zipCode int
}

// person struct is the receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	// var alex person         // assign zero value to all properties
	// fmt.Printf("%+v", alex) // struct of empty strings
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim // &variable -> give the memory address that this variable is pointing at
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy") // shortcut in golang so that you don't have to create a new pointer, golang will do it automatically if your receiver is a pointer
	jim.print()
}

// IMPORTANT!!!
// Golang is a pass-by-value language -> pass copies in receivers and arguments
// Turn "address" into "value" with "*address"
// Turn "value" into "address with "&value
// Arrays -> primitive data structure, cannot be resized
// Slices -> can grow or shrink, contain a pointer to the head of the array
// Value type -> int, float, string, bool, struct (use pointers if you want to change the underlying value)
// Reference type -> slices, maps, channels, pointers, functions (don't have to worry about pointers)

// receiver is a pointer -> "*person" means a pointer to a "person"
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer -> give the value that this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}
