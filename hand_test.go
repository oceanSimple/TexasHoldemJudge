package texas_holdem

import (
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

	//fmt.Println(cards)

	combinations := getCombinations(cards, 5)
	//for _, combination := range combinations {
	//	for _, card := range combination {
	//		print(card.String() + " ")
	//	}
	//	println()
	//}

	assert.Equal(t, len(combinations), 21)
}

func TestGetTypeAndRanks(t *testing.T) {
	var cards = getCards([]string{"♠A", "♠K", "♠Q", "♠J", "♠10"})
	var types, _ = getTypeAndRanks(cards)
	assert.Equal(t, ROYAL_FLUSH, types)

	cards = getCards([]string{"♠K", "♠Q", "♠J", "♠10", "♠9"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, STRAIGHT_FLUSH, types)

	cards = getCards([]string{"♠3", "♠2", "♥2", "♦2", "♣2"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, FOUR_OF_A_KIND, types)

	cards = getCards([]string{"♠3", "♣3", "♠2", "♥2", "♦2"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, FULL_HOUSE, types)

	cards = getCards([]string{"♠A", "♠Q", "♠J", "♠10", "♠4"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, FLUSH, types)

	cards = getCards([]string{"♠A", "♣K", "♠Q", "♠J", "♠10"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, STRAIGHT, types)

	cards = getCards([]string{"♠4", "♣3", "♠2", "♥2", "♦2"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, THREE_OF_A_KIND, types)

	cards = getCards([]string{"♠4", "♣3", "♦3", "♠2", "♥2"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, TWO_PAIR, types)

	cards = getCards([]string{"♥5", "♠4", "♣3", "♦3", "♠2"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, ONE_PAIR, types)

	cards = getCards([]string{"♠A", "♣K", "♠Q", "♠J", "♠8"})
	types, _ = getTypeAndRanks(cards)
	assert.Equal(t, HIGH_CARD, types)
}

func TestNewHand(t *testing.T) {
	var cards = getCards([]string{"♠A", "♠Q", "♠K", "♠10", "♥4", "♠J", "♦2"})
	var hand = NewHand(cards)
	assert.Equal(t, ROYAL_FLUSH, hand.Type)
	assert.Equal(t, []int{14}, hand.Ranks)

	cards = getCards([]string{"♠2", "♠3", "♠6", "♠10", "♠A", "♠Q", "♠K"})
	hand = NewHand(cards)
	assert.Equal(t, FLUSH, hand.Type)
	assert.Equal(t, []int{14, 13, 12, 10, 6}, hand.Ranks)
}
