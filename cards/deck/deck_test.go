package deck

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d.Cards) != 52 {
		t.Error("The deck length should be 52 cards")
	}
}

func TestDeck_Shuffle(t *testing.T) {
	d1 := *NewDeck()
	d2 := *NewDeck()

	if !reflect.DeepEqual(d1.Cards, d2.Cards) {
		t.Error("Both decks should be identical")
	}

	d1.Shuffle()

	if reflect.DeepEqual(d1.Cards, d2.Cards) {
		t.Error("The cards have not been shuffled")
	}
}
