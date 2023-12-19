package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// io.Copy(os.Stdout, resp.Body)
	// first thing - something that implements the writer interface
	// 	os.Stdout
	// 	value of type File
	// 	File has a fcn called 'Write'
	// 	Therefore, it implements the 'Write' interface

	// second thing - something that implements the Reader interface
	// 	resp.Body
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this number of bytes", len(bs))
	return len(bs), nil
}
