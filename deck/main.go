package main

import (
	"fmt"
	"github.com/jkuma/gophercises/deck/card"
)

func main() {
	c := card.NewDeck()
	c.Shuffle()
	c.Sort(card.DefaultSort)
	fmt.Println(*c)
}
