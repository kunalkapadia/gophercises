package main

import (
	"fmt"
	"strings"

	"github.com/kunalkapadia/gophercises/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for i := range h {
		if h[i].Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var player, dealer Hand
	var card deck.Card
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = Draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println(player.String(), ": Score:", player.MinScore())
		fmt.Println(dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = Draw(cards)
			player = append(player, card)
		}
	}
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = Draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("===FINAL CARDS===")
	fmt.Println("Player: ", player, "\nScore:", pScore)
	fmt.Println("Dealer: ", dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted!")
	case dScore > 21:
		fmt.Println("Deader busted!")
	case pScore > dScore:
		fmt.Println("You win!")
	case pScore < dScore:
		fmt.Println("You lose!")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}
