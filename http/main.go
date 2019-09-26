package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Println(io.Copy(os.Stdout, resp.Body))
	//printResponseBody(resp)
}

func printResponseBody(resp *http.Response) {
	b := make([]byte, 1000)

	for {
		n, err := resp.Body.Read(b)
		fmt.Printf("%q", b[:n])

		if err == io.EOF {
			os.Exit(1)
		}
	}
}
