//go:generate stringer -type=Suit,Rank

package card

import (
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	_ Suit = iota
	Spade
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
	maxSuit = Joker
)

type Card struct {
	Rank
	Suit
}

type Cards []Card

func NewDeck() *Cards {
	var cards Cards
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	return &cards
}

func (c *Cards) Sort(less func(c Cards) func(i, j int) bool) {
	sort.Slice(*c, less(*c))
}

func (c *Cards) Shuffle() {
	cards := make(Cards, len(*c))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i, j := range r.Perm(len(*c)) {
		cards[i] = (*c)[j]
	}

	*c = cards
}

func DefaultSort(c Cards) func(i, j int) bool {
	f := func(card Card) int {
		return int(maxRank) * int(card.Suit) + int(card.Rank)
	}

	return func(i, j int) bool {
		return f(c[i]) < f(c[j])
	}
}

func RankSort(c Cards) func(i, j int) bool {
	f := func(card Card) int {
		return int(maxSuit) * int(card.Rank) + int(card.Suit)
	}

	return func(i, j int) bool {
		return f(c[i]) < f(c[j])
	}
}