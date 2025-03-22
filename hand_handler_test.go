package texas_holdem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsFlush(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var cards = getCards([]string{"♠A", "♠J", "♠10", "♠K", "♠2"})
	ok, numArr = handler.isFlush(cards)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{14, 11, 10, 13, 2}, numArr)

	cards = getCards([]string{"♠A", "♥J", "♠10", "♠K", "♠2"})
	ok, _ = handler.isFlush(cards)
	assert.Equal(t, false, ok)
}

func TestIsStraight(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var values = []int{14, 5, 4, 3, 2}
	ok, numArr = handler.isStraight(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{5}, numArr)

	values = []int{7, 6, 5, 4, 3}
	ok, _ = handler.isStraight(values)
	assert.Equal(t, true, ok)

	values = []int{14, 6, 4, 3, 2}
	ok, _ = handler.isStraight(values)
	assert.Equal(t, false, ok)
}

func TestIsFourOfAKind(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var values = []int{14, 14, 14, 14, 2}
	ok, numArr = handler.isFourOfAKind(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{14}, numArr)

	values = []int{14, 14, 14, 2, 2}
	ok, _ = handler.isFourOfAKind(values)
	assert.Equal(t, false, ok)
}

func TestIsFullHouse(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var values = []int{14, 14, 14, 2, 2}
	ok, numArr = handler.isFullHouse(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{14, 2}, numArr)

	values = []int{14, 14, 2, 2, 2}
	ok, numArr = handler.isFullHouse(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{2, 14}, numArr)

	values = []int{14, 14, 14, 2, 3}
	ok, _ = handler.isFullHouse(values)
	assert.Equal(t, false, ok)
}

func TestIsThreeOfAKind(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var values = []int{14, 14, 14, 2, 3}
	ok, numArr = handler.isThreeOfAKind(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{14}, numArr)

	values = []int{14, 14, 2, 2, 3}
	ok, _ = handler.isThreeOfAKind(values)
	assert.Equal(t, false, ok)
}

func TestIsTwoPair(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var values = []int{14, 14, 3, 2, 2}
	ok, numArr = handler.isTwoPair(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{14, 2, 3}, numArr)

	values = []int{14, 14, 3, 3, 2}
	ok, _ = handler.isTwoPair(values)
	assert.Equal(t, true, ok)

	values = []int{14, 10, 10, 3, 3}
	ok, numArr = handler.isTwoPair(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{10, 3, 14}, numArr)

	values = []int{14, 14, 2, 3, 4}
	ok, _ = handler.isTwoPair(values)
	assert.Equal(t, false, ok)
}

func TestIsOnePair(t *testing.T) {
	var handler = HandHandler{}
	var ok bool
	var numArr []int

	var values = []int{14, 14, 4, 3, 2}
	ok, numArr = handler.isOnePair(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{14, 4, 3, 2}, numArr)

	values = []int{14, 13, 4, 4, 2}
	ok, numArr = handler.isOnePair(values)
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{4, 14, 13, 2}, numArr)

	values = []int{14, 5, 4, 3, 2}
	ok, _ = handler.isOnePair(values)
	assert.Equal(t, false, ok)
}
