package texas_holdem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCard(t *testing.T) {
	var card = NewCard(ColorMap[COLOR_SPADE], NumberMap[NUMBER_ACE])
	assert.Equal(t, "♠A", card.String())

	var card2 = NewCard(ColorMap[COLOR_HEART], NumberMap[NUMBER_JACK])
	assert.Equal(t, "♥J", card2.String())
}
