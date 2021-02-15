package main

import (
	"fmt"

	"github.com/movaua/gophercises-deck/pkg/deck"
)

func main() {
	fmt.Println("deck")

	d := deck.New(deck.WithJokers(3), deck.Shuffle)
	for _, c := range d {
		fmt.Printf("%+v\n", c)
	}
}
