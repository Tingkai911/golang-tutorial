package main

import (
	"fmt"
	"net/http"
	"time"
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

	// channel to communicate between different child go routines (send messages between routines)
	// channel are types, send string messages between different channels
	c := make(chan string)

	for _, link := range links {
		// "go" keyword = run this function in a new go routine (child go routine - possible for main routine to exit before child go routine gets competed)
		go checkLink(link, c)
	}

	// loop forever
	// for {
	// 	// beacause "<-c" is blocking, there will only be 5 request call at any point in time
	// 	go checkLink(<-c, c) // repeat the routine using the link received from the channel
	// }
	// loop forever alternative syntax (loop through the channel forever)
	for l := range c {
		// anonymous function a.k.a function literal
		go func(link string) {
			time.Sleep(5 * time.Second) // Blocking call
			checkLink(link, c)          // IMPORTANT!!! Never reference the same variable in different go routines
		}(l) // invoke the anonymous function (IMPORTANT!!! l is pass by value)
	}
}

// channel <- message (send a message to the channel)
// variable <- channel (receive a message from the channel and store it in a variable)
func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking Call
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // send the link to the channel
		return
	}
	fmt.Println(link, "is up!")
	c <- link // send the link to the channel
}
