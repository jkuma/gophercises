package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
		"http://www..org",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for u := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkUrl(url, c)
		}(u)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)

	if err == nil {
		fmt.Println(url, "is up")
		c <- url
		return
	}

	fmt.Println(url, "is down")
	c <- url
}
