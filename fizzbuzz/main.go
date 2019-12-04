package main

import (
	"fmt"
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"os"
)

func main() {
	must(database.Init("./fizzbuzz.db"))
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
