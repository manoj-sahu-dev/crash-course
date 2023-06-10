package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, error := http.Get("http://www.google.com")

	if error != nil {
		fmt.Println("error getting url: " + error.Error())
		os.Exit(1)
	}
	fmt.Println(resp.Body)

	// byteSlice := make([]byte, 1024*32)
	// resp.Body.Read(byteSlice)
	// fmt.Println("--------------------------------")
	// fmt.Println(string(byteSlice))

	io.Copy(os.Stdout, resp.Body)
}
