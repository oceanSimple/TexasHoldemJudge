package texas_holdem

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCombinations(t *testing.T) {
	cards := []Card{
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_ACE]),
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_KING]),
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_QUEEN]),
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_JACK]),
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_10]),
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_9]),
		NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_8]),
	}

	fmt.Println(cards)

	combinations := getCombinations(cards, 5)
	for _, combination := range combinations {
		for _, card := range combination {
			print(card.String() + " ")
		}
		println()
	}

	assert.Equal(t, len(combinations), 21)
}
