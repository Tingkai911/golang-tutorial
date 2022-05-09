package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// deferExample()
	// deferExample2()
	// deferExample3()
	deferExample4()
}

func deferExample() {
	fmt.Println("start")
	defer fmt.Print("middle") // execute this after the function exits, but before the function returns
	fmt.Print("end")
}

// "defer" executes in LIFO order
func deferExample2() {
	defer fmt.Println("start")
	defer fmt.Print("middle")
	defer fmt.Print("end")
}

// "defer" is normally used to close resources
func deferExample3() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer res.Body.Close() // "defer" closing this body until the end of the functions

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error", err)
	}
	fmt.Printf("Response: %s", string(robots))
}

func deferExample4() {
	a := "start"
	defer fmt.Println(a) // "start" will be printed beacuse a = "start" at the time "defer" is used
	a = "end"
}
