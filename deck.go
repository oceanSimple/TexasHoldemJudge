package texas_holdem

import (
	"fmt"
	"math/rand"
	"strings"
)

type Deck struct {
	Deck []Card
}

// Obtain a full deck of cards (excluding kings)
func NewDeck() Deck {
	var deck = Deck{
		Deck: obtainNewDeck(),
	}

	deck.Sort()

	return deck
}

// Sort all the cards, in the  order of ♠ ♥ ♦ ♣
func (deck *Deck) Sort() {
	SortCard(deck.Deck)
}

// Shuffle the deck
func (deck *Deck) Shuffle() {
	for i := len(deck.Deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck.Deck[i], deck.Deck[j] = deck.Deck[j], deck.Deck[i]
	}
}

func (deck *Deck) String() string {
	var builder = strings.Builder{}

	for i := 0; i < len(deck.Deck); i++ {
		builder.WriteString(fmt.Sprintf("%-4s", deck.Deck[i].String()))
		if i == 12 || i == 25 || i == 38 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

// Obtain a full deck of cards (excluding kings)
func obtainNewDeck() []Card {
	var deck []Card
	for _, color := range ColorMap {
		for _, number := range NumberMap {
			deck = append(deck, NewCard(color, number))
		}
	}
	return deck
}
