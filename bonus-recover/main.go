package main

import "fmt"

func main() {
	// recoverExample1()
	// recoverExample2()
	recoverExample3()
}

func recoverExample1() {
	fmt.Println("start")
	// defer keyword takes a function call
	defer func() {
		// recover() returns nil if the function is not panicking
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}() // need to invoke the functiom
	panic("something bad happened")
	fmt.Println("end")
}

// functions higher up the call stack will still continue despite paniker() has panicked
func recoverExample2() {
	fmt.Println("start")
	panicker()
	fmt.Println("end") // will still execute
}

// function that panics will stop execution after the panic
func panicker() {
	fmt.Println("about to panic")
	// defer keyword takes a function call
	defer func() {
		// recover() returns nil if the function is not panicking
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}() // need to invoke the functiom
	panic("something bad happened")
	fmt.Println("done panicking")
}

func recoverExample3() {
	fmt.Println("start")
	panicker2()
	fmt.Println("end") // will not execute
}

func panicker2() {
	fmt.Println("about to panic")
	// defer keyword takes a function call
	defer func() {
		// recover() returns nil if the function is not panicking
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			panic(err) // rethrow the panic
		}
	}() // need to invoke the functiom
	panic("something bad happened")
	fmt.Println("done panicking")
}
