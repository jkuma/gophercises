package deck

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Unit  string
	Color string
}

func NewDeck() *Deck {
	var deck Deck

	units := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	colors := []string{"spade", "diamond", "heart", "club"}

	for _, c := range colors {
		for _, u := range units {
			deck.Cards = append(deck.Cards, Card{Unit: u, Color: c})
		}
	}

	return &deck
}

func (d *Deck) Shuffle() {
	var cards []Card

	rand.Seed(time.Now().UnixNano())

	for _ = range d.Cards {
		i := rand.Intn(len(d.Cards))
		cards = append(cards, d.Cards[i])
		d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
	}

	d.Cards = cards
}
