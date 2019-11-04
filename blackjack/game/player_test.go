package game

import (
	"fmt"
	"github.com/jkuma/gophercises/deck/card"
	"strings"
	"testing"
)

func TestNewPlayers(t *testing.T) {
	players := NewPlayers(2)

	if len(players) != 2 {
		t.Error("The number of players should be 2")
	}
}

func TestPlayer_CardsScore(t *testing.T) {
	p := Player{
		Name: PlayerName,
		Cards: card.Cards{
			card.Card{Rank: card.Three, Suit: card.Spade},
			card.Card{Rank: card.Four, Suit: card.Spade},
		},
	}

	if p.CardsScore() != 7 {
		t.Error("The player score is not correct, it should be 7")
	}
}

func TestPlayer_DrawCard(t *testing.T) {
	want := card.Card{Rank: card.Two, Suit: card.Spade}

	deck := card.Cards{want}
	p := Player{Name: PlayerName}

	c := p.DrawCard(&deck)

	if c != want {
		t.Error("The card should be Two of spades")
	}
}

func TestPlayer_Hand(t *testing.T) {
	want := []string{
		fmt.Sprint(card.Card{Rank: card.Three, Suit: card.Spade}),
		fmt.Sprint(card.Card{Rank: card.Four, Suit: card.Spade}),
	}

	p := Player{
		Name: PlayerName,
		Cards: card.Cards{
			card.Card{Rank: card.Three, Suit: card.Spade},
			card.Card{Rank: card.Four, Suit: card.Spade},
		},
	}

	if p.Hand() != strings.Join(want, ", ") {
		t.Error("Player hand is not correctly printed")
	}
}
