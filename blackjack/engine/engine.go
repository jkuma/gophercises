package engine

import (
	"fmt"
	"github.com/jkuma/gophercises/deck/card"
	"strings"
)

const (
	PlayerName = "Player"
	DealerName = "Dealer"
)

type Draw interface {
	DrawCard(d *card.Cards) card.Card
}

type Player struct {
	Name string
	Cards card.Cards
	Score int8
	Dealer bool
}

func Run() {
	// Init deck
	d := card.NewDeck()
	d.Sort(card.DefaultSort)
	d.Shuffle()

	// Init dealer and player.
	players := NewPlayers(2)

	// Deal two cards to each player from deck.
	fmt.Println("...First Round...")
	firstRound(players, d)

	// Hit or stand??
	fmt.Println("...Hit or Stand???...")
	secondRound(players, d)

	for _, p := range players {
		fmt.Println(p.Name, p.CardsScore())
	}
}

func firstRound(players []Player, d card.Cards) {
	for i := 0; i < 2; i++ {
		for k, p := range players {
			c := p.DrawCard(d)
			players[k].Cards = append(players[k].Cards, c)

			if i == 1 && p.Dealer {
				fmt.Println(p.Name, "drawn: hidden card")
				continue
			}

			fmt.Println(p.Name, "has drawn:", c)
		}
	}
}

func secondRound(players []Player, d card.Cards) {
	for k, p := range players {
		var i string
		if p.Dealer {continue}

		for i != "s" {
			fmt.Println(p.Name, "hand is", p.Hand(), "\nDo you want to (h)it or (s)tand?")
			fmt.Scanf("%s", &i)
			if i != "h" {break}

			c := p.DrawCard(d)
			players[k].Cards = append(players[k].Cards, c)
			fmt.Println(p.Name, "draw", c)

			if p.CardsScore() > Blackjack {
				fmt.Println(p.Name, "loses, hand is", p.CardsScore())
				break
			}
		}
	}
}

func NewPlayers(n int) []Player {
	players := make([]Player, n)

	for j, _ := range players {
		players[j].Name = fmt.Sprintf("%v%v", PlayerName, j)
	}

	return append(players, Player{Dealer: true, Name: DealerName})
}

func (p *Player) DrawCard(d card.Cards) card.Card {
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
	return strings.Join(s, ", ")
}
