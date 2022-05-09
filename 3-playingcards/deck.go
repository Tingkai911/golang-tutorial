package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck" which is a slice of strigs
type deck []string

// Function that returns a string
// No receiver needed
func newDeck() deck {
	// <new variable> <variable name> <only value of this type will be assigned> = "Ace of Spades"
	// var card string = "Ace of Spades"
	// Shorter way of writing the same thing, rely on the go complier to create type for us
	// Only use ":=" when defining a new variable, use  "=" for reassignment
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// IMPORTANT! - append function returns a new slice instead of modifiying the current slice
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function - "deck" is the receiver
// Any variable of type "deck" get access to the print method
func (d deck) print() {
	// <index>, <current card> := range <slice to iterate>
	// ":=" is used beacuse we are re-initialising i and card for each loop
	// both i and card must be used in the loop or else there will be compliation errors
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function that returns two separate values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Type conversion -> <type we want>(<type we are providing>)
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// Convert string to byte slice
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// If nothing went wrong, err == nil
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// Set a random source/seed to rand
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swap positions in the array
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
