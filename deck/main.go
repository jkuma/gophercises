package main

import (
	"fmt"
	"github.com/jkuma/gophercises/deck/card"
)

func main() {
	c := card.NewDeck()
	c.Sort(card.DefaultSort)
	fmt.Println(c)
	
	filters := card.Cards{
		card.Card{Rank: card.Ace, Suit: card.Heart},
		card.Card{Rank: card.Ace, Suit: card.Club},
		card.Card{Rank: card.Ace, Suit: card.Spade},
		card.Card{Rank: card.Ace, Suit: card.Diamond},
	}

	c.Filter(filters)
	fmt.Println(c)

	c.Shuffle()
	fmt.Println(c)
}
