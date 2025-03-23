package texas_holdem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeck(t *testing.T) {
	//var deck = NewDeck()
	//fmt.Println(deck.String())
}

func TestDeck_Shuffle(t *testing.T) {
	//var deck = NewDeck()
	//fmt.Println(deck.String())
	//
	//fmt.Println("Shuffling the deck...")
	//
	//deck.Shuffle()
	//fmt.Println(deck.String())
}

func TestDeck_Draw(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()

	var drawnCards = deck.Draw(5)
	assert.Equal(t, 5, len(drawnCards))
	assert.Equal(t, 47, len(deck.Deck))
}
