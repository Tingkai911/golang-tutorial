package main

import "fmt"

// anything type(receiver) with methods that matches all the methods inside this interface is now an honorary member of type "bot"
type bot interface {
	getGretting() string
}

type englishBot struct{}
type spanishBot struct{}

// Concrete Type (can create a value directly) -> map, struct, int, string, englishBot
// Interface Type (cannot create a value directly) -> bot

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// since "englishBot" and "spanishBot" is an honorary member of "bot", they can use printGreeting
func printGreeting(b bot) {
	fmt.Println(b.getGretting())
}

// "englishBot" is an honorary member of type "bot"
func (englishBot) getGretting() string {
	// VERY custom logic
	return "Hi There!"
}

// "spanishBot" is an honorary member of type "bot"
func (spanishBot) getGretting() string {
	// VERY custom logic
	return "Hola!"
}
