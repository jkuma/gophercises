package card

import "testing"

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Error("A new deck should contain 52 cards.")
	}
}
