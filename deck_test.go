package texas_holdem

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	var deck = NewDeck()
	fmt.Println(deck.String())
}

func TestDeck_Shuffle(t *testing.T) {
	var deck = NewDeck()
	fmt.Println(deck.String())

	fmt.Println("Shuffling the deck...")

	deck.Shuffle()
	fmt.Println(deck.String())
}
