package deck

// Card represents a playing card
type Card struct {
	Suit
	Rank
}

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
