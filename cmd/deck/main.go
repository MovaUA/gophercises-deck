package main

import (
	"fmt"

	"github.com/movaua/gophercises-deck/pkg/deck"
)

func main() {
	fmt.Println("deck")

	d := deck.New(deck.WithJoker(), deck.Shuffle)
	for _, c := range d {
		fmt.Printf("%+v\n", c)
	}
}
