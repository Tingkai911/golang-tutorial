package main

import (
	"fmt"
	"net/http"
)

// golang does not have exceptions
func main() {
	// panicExample1()
	// panicExample2()
	// panicExample3()
	panicExample4()
}

func panicExample1() {
	a, b := 1, 0
	ans := a / b // will generate a panic
	fmt.Println(ans)
}

func panicExample2() {
	fmt.Println("start")
	panic("something bad happened") // throw your own panic
	fmt.Println("end")
}

func panicExample3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil) // go to localhost:8080

	// will have error if port 8080 is already used
	if err != nil {
		panic(err.Error())
	}
}

// defer will always run before a panic
func panicExample4() {
	fmt.Println("start")
	defer fmt.Print("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}
