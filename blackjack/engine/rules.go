package engine

import "github.com/jkuma/gophercises/deck/card"

const (
	Blackjack = 21
)

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
