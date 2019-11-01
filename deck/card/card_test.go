package card

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := NewDeck()

	if len(*cards) != 52 {
		t.Error("A new deck should contain 52 cards.")
	}
}

func TestDecks(t *testing.T) {
	decks := Decks(2)

	if len(decks) != 2 {
		t.Error("It should contains two decks.")
	}
}

func TestCards_Filter(t *testing.T) {
	filters := Cards{
		Card{Rank: Ace, Suit: Heart},
		Card{Rank: Ace, Suit: Club},
		Card{Rank: Ace, Suit: Spade},
		Card{Rank: Ace, Suit: Diamond},
	}

	d := NewDeck()
	d.Filter(filters)

	for _, c := range *d {
		for _, f := range filters {
			if c == f {
				t.Errorf("The deck should not contain %v", f)
			}
		}
	}
}

func TestCards_Shuffle(t *testing.T) {
	d1 := NewDeck()
	d2 := NewDeck()
	d2.Shuffle()

	if reflect.DeepEqual(d1, d2) {
		t.Error("The deck shuffling don't work")
	}
}
