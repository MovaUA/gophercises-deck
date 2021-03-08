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
	Suit `json:"suit,omitempty"`
	Rank `json:"rank,omitempty"`
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// New creates Card slice (a new deck of playing cards)
func New(opts ...OptFunc) []Card {
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

// OptFunc is an option function type
type OptFunc func([]Card) []Card

// WithJokers returns an option which adds specified number of Jokers
func WithJokers(n int) OptFunc {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

// Shuffle returns shuffled slice of cards
func Shuffle(cards []Card) []Card {
	seed := time.Now().UnixNano()
	return shuffle(cards, seed)
}

// shuffle returns shuffled slice of cards based on the provided seed
func shuffle(cards []Card, seed int64) []Card {
	rnd := rand.New(rand.NewSource(seed))
	for i, j := range rnd.Perm(len(cards)) {
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

// WithDefaultSort returns an options which sorts cards with default sort
func WithDefaultSort() OptFunc {
	return func(cards []Card) []Card {
		sort.SliceStable(cards, Less(cards))
		return cards
	}
}

// WithSort returns an option which sorts cards with provided less function
func WithSort(less func([]Card) func(i, j int) bool) OptFunc {
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

// FilterFunc is filter function, which returns true if the provided card must be selected,
// and returns false if the card must be filtered out.
type FilterFunc func(card Card) bool

// WithFilter is an option which provides a filter func
func WithFilter(f FilterFunc) OptFunc {
	return func(cards []Card) []Card {
		filtered := make([]Card, 0, len(cards))
		for _, card := range cards {
			if f(card) {
				filtered = append(filtered, card)
			}
		}
		return filtered
	}
}
