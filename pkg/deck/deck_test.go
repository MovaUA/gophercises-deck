package deck

import "fmt"

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
