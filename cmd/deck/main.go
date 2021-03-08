package main

import (
	"encoding/json"
	"os"

	"github.com/movaua/gophercises-deck/pkg/deck"
)

func main() {
	cards := deck.New(deck.WithJokers(3), deck.Shuffle)

	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.SetIndent("", "  ")

	jsonEncoder.Encode(cards)
}
