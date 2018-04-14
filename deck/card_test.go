package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Spade, Rank: Two})
	fmt.Println(Card{Suit: Diamond, Rank: Six})
	fmt.Println(Card{Suit: Club, Rank: Nine})
	fmt.Println(Card{Suit: Heart, Rank: Queen})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Two of Spades
	// Six of Diamonds
	// Nine of Clubs
	// Queen of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 4*13 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	firstCard := Card{Suit: Spade, Rank: Ace}
	if cards[0] != firstCard {
		t.Error("Expected first card to be Ace of Spades")
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	firstCard := Card{Suit: Spade, Rank: Ace}
	if cards[0] != firstCard {
		t.Error("Expected first card to be Ace of Spades")
	}
}

func TestJokers(t *testing.T) {
	nJokers := 5
	cards := New(Jokers(nJokers))
	jokerCount := 0
	for _, card := range cards {
		if card.Suit == Joker {
			jokerCount++
		}
	}

	if jokerCount != nJokers {
		t.Errorf("Expected %d jokers, received %d", nJokers, jokerCount)
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(5))
	// 13 ranks * 4 suits * 5 decks
	if len(cards) != 13*4*5 {
		t.Errorf("Expected %d number of cards, got %d", 13*4*5, len(cards))
	}
}
