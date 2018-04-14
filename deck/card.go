//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // special case
)

const (
	MinSuit = Spade
	MaxSuit = Heart
)

type Rank uint8

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
	MinRank = Ace
	MaxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func (c Card) absRank() int {
	return int(c.Suit)*int(MaxRank) + int(c.Rank)
}

// New func returns a new deck of cards.
// It accepts an optional number of function which accepts and returns slice of cards.
func New(opts ...func(cards []Card) []Card) []Card {
	var cards []Card
	for suit := MinSuit; suit <= MaxSuit; suit++ {
		for rank := MinRank; rank <= MaxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	// Run cards over each of opts
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].absRank() < cards[j].absRank()
	}
}

func Shuffle(cards []Card) []Card {
	shuffledCards := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, j := range r.Perm(len(cards)) {
		shuffledCards[i] = cards[j]
	}
	return shuffledCards
}

func Jokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

func Filter(f func(c Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}

		return ret
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
