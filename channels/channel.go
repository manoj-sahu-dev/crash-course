package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://github.com/",
		"https://facebook.com/",
		"https://google.com/",
		"https://amazon.in/",
		"https://golang.org/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		// go checkLink(l, c)
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(li, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(time.Second)
	_, error := http.Get(link)
	if error != nil {
		c <- link
	}
	fmt.Println(link, "is up!")
	c <- link
}
