package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("You must provide a text file full path.")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	_,_ = io.Copy(os.Stdout, file)
	os.Exit(1)
}
