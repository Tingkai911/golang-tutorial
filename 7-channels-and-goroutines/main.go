package main

import (
	"fmt"
	"net/http"
)

// main will run in its own go routine (main routine - control when the program exits)
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Bad! Check the links one at a time! Takes too long if you are checking many links!
	// for _, link := range links {
	// 	checkLink(link)
	// }

	// channel to communicate between different child go routines (send messages between routines)
	// channel are types, send string messages between different channels
	c := make(chan string)

	for _, link := range links {
		// "go" keyword = run this function in a new go routine (child go routine - possible for main routine to exit before child go routine gets competed)
		go checkLink(link, c)
	}

	// the whole loop is not executed instantaneously due to blocking statement
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // Blocking Call, code stuck here while waiting for a response from a channel
	}
}

// channel <- message (send a message to the channel)
// variable <- channel (receive a message from the channel and store it in a variable)
func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking Call
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
