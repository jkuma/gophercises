package game

import (
	"fmt"
	"github.com/jkuma/gophercises/deck/card"
)

type Game struct {
	Deck    *card.Cards
	Dealer  *Player
	Players []Player
}

func (g Game) DealerDraw() {
	for g.Dealer.CardsScore() < 17 {
		c := g.Dealer.DrawCard(g.Deck)
		fmt.Println(g.Dealer.Name, "has drawn", c)
	}
}

func Run() {
	// Init deck
	d := card.NewDeck()
	d.Sort(card.DefaultSort)
	d.Shuffle()

	// Init Game
	g := Game{
		Deck:    &d,
		Players: NewPlayers(2),
		Dealer:  &Player{Name: DealerName, Dealer: true},
	}

	// Deal two cards to each player from deck.
	fmt.Println("...First Round...")
	firstRound(g)

	// Hit or stand??
	fmt.Println("...Hit or Stand???...")
	secondRound(g)

	// Score
	fmt.Println("...Results...")
	finalRound(g)
}

func firstRound(g Game) {
	for i := 0; i < 2; i++ {
		for k, _ := range g.Players {
			g.Players[k].DrawCard(g.Deck)
		}

		g.Dealer.DrawCard(g.Deck)
	}
}

func secondRound(g Game) {
	for k, p := range g.Players {
		var i string

		for i != "s" {
			fmt.Println(g.Dealer.Name, "hand is", g.Dealer.Hand())
			fmt.Println(p.Name, "hand is", p.Hand())
			fmt.Println("Do you want to (h)it or (s)tand?")
			fmt.Scanf("%s", &i)
			if i != "h" {
				break
			}

			c := p.DrawCard(g.Deck)
			g.Players[k].Cards = append(g.Players[k].Cards, c)
			fmt.Println(p.Name, "draw", c)

			if p.CardsScore() > Blackjack {
				fmt.Println(p.Name, "loses, hand is", p.CardsScore())
				break
			}
		}
	}

	g.DealerDraw()
	fmt.Println(g.Dealer.Name, "hand is", g.Dealer.Cards)
}

func finalRound(g Game) {
	ds := g.Dealer.CardsScore()
	for _, p := range g.Players {
		ps := p.CardsScore()
		switch {
		case ps > Blackjack:
			fmt.Println(p.Name, "busted")
			p.Score++
		case ds > Blackjack:
			fmt.Println("Dealer busted")
		case ps > ds:
			fmt.Println(p.Name, "wins against the bank")
			p.Score++
		case ds > ps:
			fmt.Println(p.Name, "loses")
		case ds == ps:
			fmt.Println("Draw")
		}
	}
}
