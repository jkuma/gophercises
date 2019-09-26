package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct {}
type frenchBot struct {}

func main() {
	eb, fb := englishBot{}, frenchBot{}

	printGreetings(eb)
	printGreetings(fb)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (eb englishBot) getGreetings() string {
	return "Greetings!!"
}


func (fb frenchBot) getGreetings() string {
	return "Bienvenue!!!"
}
