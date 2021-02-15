//go:generate stringer -type=Suit
//go:generate stringer -type=Rank

// Package deck is a general interface for the deck of the playing cards
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit of the playing card
type Suit uint8

// All suits of the playing card
const (
	Club Suit = iota
	Diamond
	Heart
	Spade
	Joker
)

var suits = [...]Suit{Club, Diamond, Heart, Spade}

// Rank of the playing card
type Rank uint8

// All ranks of the playing card
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
)

// Card represents a playing card
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// New creates Card slice (a new deck of playing cards)
func New(opts ...func([]Card) []Card) []Card {
	cards := make([]Card, 0, 4*13)
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// WithJoker returns an option which adds Joker
func WithJoker() func([]Card) []Card {
	return func(cards []Card) []Card {
		return append(cards, Card{Suit: Joker})
	}
}

// Shuffle returns shuffled slice of cards
func Shuffle(cards []Card) []Card {
	return WithShuffle()(cards)
}

// WithShuffle returns an option which shuffles a slice of cards
func WithShuffle() func([]Card) []Card {
	return func(cards []Card) []Card {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i, j := range rnd.Perm(len(cards)) {
			cards[i], cards[j] = cards[j], cards[i]
		}
		return cards
	}
}

// WithDefaultSort returns an options which sorts cards with default sort
func WithDefaultSort() func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.SliceStable(cards, Less(cards))
		return cards
	}
}

// WithSort returns an option which sorts cards with provided less function
func WithSort(less func([]Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.SliceStable(cards, less(cards))
		return cards
	}
}

// Less returns default soft less function which sorts first by Suit, then by Rank
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return abs(cards[i]) < abs(cards[j])
	}
}

func abs(card Card) int {
	return int(card.Suit)*(int(maxRank)+1) + int(card.Rank)
}
