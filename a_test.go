package texas_holdem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetACard(t *testing.T) {
	spadeJ := GetACard(COLOR_SPADE, NUMBER_JACK)
	assert.Equal(t, "♠J", spadeJ.String())

	heart10 := GetACard(COLOR_HEART, NUMBER_10)
	assert.Equal(t, "♥10", heart10.String())

	diamond3 := GetACard(COLOR_DIAMOND, NUMBER_3)
	assert.Equal(t, "♦3", diamond3.String())

	club7 := GetACard(COLOR_CLUB, NUMBER_7)
	assert.Equal(t, "♣7", club7.String())
}

func TestGetColor(t *testing.T) {
	spade := GetColor(COLOR_SPADE)
	assert.Equal(t, 1, spade.Code)

	heart := GetColor(COLOR_HEART)
	assert.Equal(t, 2, heart.Code)

	diamond := GetColor(COLOR_DIAMOND)
	assert.Equal(t, 3, diamond.Code)

	club := GetColor(COLOR_CLUB)
	assert.Equal(t, 4, club.Code)
}

func TestGetNumber(t *testing.T) {
	ace := GetNumber(NUMBER_ACE)
	assert.Equal(t, 1, ace.Code)

	two := GetNumber(NUMBER_2)
	assert.Equal(t, 2, two.Code)

	three := GetNumber(NUMBER_3)
	assert.Equal(t, 3, three.Code)

	four := GetNumber(NUMBER_4)
	assert.Equal(t, 4, four.Code)

	five := GetNumber(NUMBER_5)
	assert.Equal(t, 5, five.Code)

	six := GetNumber(NUMBER_6)
	assert.Equal(t, 6, six.Code)

	seven := GetNumber(NUMBER_7)
	assert.Equal(t, 7, seven.Code)

	eight := GetNumber(NUMBER_8)
	assert.Equal(t, 8, eight.Code)

	nine := GetNumber(NUMBER_9)
	assert.Equal(t, 9, nine.Code)

	ten := GetNumber(NUMBER_10)
	assert.Equal(t, 10, ten.Code)

	jack := GetNumber(NUMBER_JACK)
	assert.Equal(t, 11, jack.Code)

	queen := GetNumber(NUMBER_QUEEN)
	assert.Equal(t, 12, queen.Code)

	king := GetNumber(NUMBER_KING)
	assert.Equal(t, 13, king.Code)
}
