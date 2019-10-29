package main

import (
	"fmt"
	deck2 "github.com/jkuma/gophercises/deck/deck"
)

func main() {
	deck := deck2.NewDeck()
	deck.Shuffle()
	fmt.Println(deck)
}
