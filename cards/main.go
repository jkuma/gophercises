package main

import (
	"fmt"
	"github.com/jkuma/gophercises/cards/deck"
)

func main() {
	d := deck.NewDeck()
	d.Shuffle()
	fmt.Println(d)
}
