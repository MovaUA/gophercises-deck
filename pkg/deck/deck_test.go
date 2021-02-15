package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Club, Rank: Ace})
	fmt.Println(Card{Suit: Heart, Rank: King})
	fmt.Println(Card{Suit: Diamond, Rank: Two})
	fmt.Println(Card{Suit: Spade, Rank: Ten})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Clubs
	// King of Hearts
	// Two of Diamonds
	// Ten of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	want := 13 * 4
	got := len(cards)
	if got != want {
		t.Errorf("Wrong number of cards: got %d, want %d", got, want)
	}
}
func TestWithJoker(t *testing.T) {
	cards := New(WithJoker(), WithJoker())
	want := 13*4 + 2
	got := len(cards)
	if got != want {
		t.Errorf("Wrong number of cards: got %d, want %d", got, want)
	}
	jokerCount := 0
	for _, card := range cards {
		if card.Suit == Joker {
			jokerCount++
		}
	}
	if jokerCount != 2 {
		t.Errorf("Wrong number of %ss: got %d, want %d", Joker, jokerCount, 2)
	}
}

func TestWithDefaultSort(t *testing.T) {
	cards := New(WithDefaultSort())
	first, last := cards[0], cards[len(cards)-1]
	wantFirst, wantLast := Card{Suit: Club, Rank: Ace}, Card{Suit: Spade, Rank: King}
	if first != wantFirst {
		t.Errorf("Wrong first card: got %q, want %q", first, wantFirst)
	}
	if last != wantLast {
		t.Errorf("Wrong last card: got %q, want %q", last, wantLast)
	}
}
