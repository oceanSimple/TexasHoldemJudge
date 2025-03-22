package texas_holdem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockCards() [][]Card {
	// ♠K ♥8 ♥10 ♣7 ♥5
	var publicCards = []Card{GetACard(COLOR_SPADE, NUMBER_KING),
		GetACard(COLOR_HEART, NUMBER_8),
		GetACard(COLOR_HEART, NUMBER_10),
		GetACard(COLOR_CLUB, NUMBER_7),
		GetACard(COLOR_HEART, NUMBER_5)}

	// ♠K ♥K ♠8 ♥8 ♥10 ♣7 ♥5
	var card0 = append(publicCards,
		[]Card{GetACard(COLOR_HEART, NUMBER_KING), GetACard(COLOR_SPADE, NUMBER_8)}...)

	// ♠K ♣K ♥8 ♣8 ♥10 ♣7 ♥5
	var card1 = append(publicCards,
		[]Card{GetACard(COLOR_CLUB, NUMBER_KING), GetACard(COLOR_CLUB, NUMBER_8)}...)

	// ♥5 ♦5 ♠K ♥J ♥10 ♥8 ♣7
	var card2 = append(publicCards,
		[]Card{GetACard(COLOR_HEART, NUMBER_JACK), GetACard(COLOR_DIAMOND, NUMBER_5)}...)

	// ♥10 ♥8 ♥5 ♥4 ♥2 ♠K ♣7
	var card3 = append(publicCards,
		[]Card{GetACard(COLOR_HEART, NUMBER_4), GetACard(COLOR_HEART, NUMBER_2)}...)

	// ♥10 ♥9 ♥8 ♥6 ♥5 ♠K ♣7
	var card4 = append(publicCards,
		[]Card{GetACard(COLOR_HEART, NUMBER_9), GetACard(COLOR_HEART, NUMBER_6)}...)

	return [][]Card{card0, card1, card2, card3, card4}
}

func TestHandler_CompareTwoCards(t *testing.T) {
	var cards = mockCards()
	var handler = Handler{handHandler: HandHandler{}}

	var flag, _ = handler.CompareTwoCards(cards[0], cards[1])
	assert.Equal(t, 0, flag)

	flag, _ = handler.CompareTwoCards(cards[0], cards[2])
	assert.Equal(t, 1, flag)

	flag, _ = handler.CompareTwoCards(cards[0], cards[3])
	assert.Equal(t, -1, flag)

	flag, _ = handler.CompareTwoCards(cards[0], cards[4])
	assert.Equal(t, -1, flag)

	flag, _ = handler.CompareTwoCards(cards[3], cards[4])
	assert.Equal(t, -1, flag)
}

func TestHandler_GetMaxCard(t *testing.T) {
	var cards = mockCards()
	var handler = Handler{handHandler: HandHandler{}}

	var indexes, _ = handler.GetMaxCard([][]Card{cards[0], cards[1]})
	assert.Equal(t, []int{0, 1}, indexes)

	indexes, _ = handler.GetMaxCard(cards)
	assert.Equal(t, []int{4}, indexes)
}
