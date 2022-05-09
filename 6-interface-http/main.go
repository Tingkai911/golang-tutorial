package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// "logWriter" is now implementing the Writer interface
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(resp) // wrong way to do it, you will not see the response body

	// Use the Reader interface to get the response body
	// bs := make([]byte, 99999) // create a empty byte size with space for 99999 elements, Reader is not set up to resize this slice so good to allocate more memory
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Build in helper function to get the response body (standard library that implements the Writer interface)
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
