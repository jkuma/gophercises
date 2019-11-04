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

func NewDeck() Cards {
	var cards Cards
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	return cards
}

func Decks(n int) []Cards {
	decks := make([]Cards, n)

	for i := 0; i < n; i++ {
		decks[i] = NewDeck()
	}

	return decks
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

func (c *Cards) Filter(cards Cards) {
	for i, card := range *c {
		for _, fcard := range cards {
			if card == fcard {
				*c = append((*c)[:i], (*c)[i+1:]...)
			}
		}
	}
}

func (c *Cards) Jokers() {
	for i:= 1; i <= 4; i++ {
		*c = append(*c, Card{Rank: Rank(i), Suit: Joker})
	}
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
