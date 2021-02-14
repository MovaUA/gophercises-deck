package deck

// Card represents a playing card
type Card struct {
	Suit  Suit
	Value int
}

// Suit is the playing card suit
type Suit int

// All suits of the playing card
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)
