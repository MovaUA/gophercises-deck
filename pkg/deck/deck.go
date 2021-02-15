//go:generate stringer -type=Suit
//go:generate stringer -type=Rank

// Package deck is a general interface for the deck of the playing cards
package deck

import "fmt"

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
