package deck

import (
	"bytes"
	_ "embed"
	"encoding/json"
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
func TestWithJokers(t *testing.T) {
	wantJokers := 3
	cards := New(WithJokers(wantJokers))
	wantLen := 13*4 + wantJokers
	gotLen := len(cards)
	if gotLen != wantLen {
		t.Errorf("Wrong number of cards: got %d, want %d", gotLen, wantLen)
	}
	gotJokers := 0
	for _, card := range cards {
		if card.Suit == Joker {
			gotJokers++
		}
	}
	if gotJokers != wantJokers {
		t.Errorf("Wrong number of %ss: got %d, want %d", Joker, gotJokers, wantJokers)
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

func TestWithFilter(t *testing.T) {
	wantRank := Ten
	f := func(card Card) bool {
		return card.Rank == wantRank
	}
	cards := New(WithFilter(f))
	for _, card := range cards {
		if card.Rank != wantRank {
			t.Errorf("Wrong card: got %q, want Rank of %q", card, wantRank)
		}
	}
}

//go:embed _tests/wantShuffle.json
var wantShuffle []byte

func TestShuffle(t *testing.T) {
	var want []Card
	{
		jsonDecoder := json.NewDecoder(bytes.NewReader(wantShuffle))
		if err := jsonDecoder.Decode(&want); err != nil {
			t.Fatal(err)
		}
	}

	testShuffle := func(cards []Card) []Card {
		return shuffle(cards, 0)
	}

	got := New(testShuffle)

	if len(got) != len(want) {
		t.Fatalf("len: want %d, got %d\n", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Fatalf("index %d: want %+v, got %+v\n", i, want[i], got[i])
		}
	}
}
