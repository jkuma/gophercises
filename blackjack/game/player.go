package game

import (
	"fmt"
	"github.com/jkuma/gophercises/deck/card"
	"strings"
)

const (
	Blackjack  = 21
	PlayerName = "Player"
	DealerName = "Dealer"
)

type Player struct {
	Name   string
	Cards  card.Cards
	Score  uint8
	Dealer bool
}

func NewPlayers(n int) []Player {
	players := make([]Player, n)

	for j, _ := range players {
		players[j].Name = fmt.Sprintf("%v%v", PlayerName, j)
	}

	return players
}

func (p *Player) DrawCard(d *card.Cards) card.Card {
	c, err := d.DealCard()

	if err != nil {
		panic(err)
	}

	p.Cards = append(p.Cards, c)

	return c
}

func (p *Player) Hand() string {
	s := make([]string, len(p.Cards))
	for i := range p.Cards {
		s[i] = fmt.Sprint(p.Cards[i])
	}

	if len(p.Cards) == 2 && p.Dealer {
		s[1] = "Hidden"
	}

	return strings.Join(s, ", ")
}

func (p *Player) CardsScore() uint8 {
	var min, max uint8

	for _, c := range p.Cards {
		switch c.Rank {
		case card.Ace:
			min += 1
			max += 11
		case card.Jack, card.Queen, card.King:
			min += 10
			max += 10
		default:
			min += uint8(c.Rank)
			max += uint8(c.Rank)
		}
	}

	if max <= Blackjack {
		return max
	}

	return min
}
