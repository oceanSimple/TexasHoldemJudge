package texas_holdem

import "errors"

type IHandler interface {
	// 这里的card是玩家的 2张手牌 + 5张公共牌 = 一共七张牌
	CompareTowHands(c1, c2 []Card) (bool, error)
}

type Handler struct {
	handHandler IHandHandler
}

func (h Handler) CompareTowHands(c1, c2 []Card) (bool, error) {
	// 首先判断是否是一共七张牌
	if len(c1) != 7 {
		return false, errors.New("c1 must be 7 cards")
	}
	if len(c2) != 7 {
		return false, errors.New("c2 must be 7 cards")
	}

	// 将 []Card 转换成 Hand
	var hand1 = h.cardsToHand(c1)
	var hand2 = h.cardsToHand(c2)

	return h.compareTowHands(hand1, hand2), nil
}

func (h Handler) cardsToHand(cards []Card) *Hand {
	hand := &Hand{
		Cards: cards,
		Type:  0,
		Ranks: make([]int, 0),
	}

	return hand
}

func (h Handler) compareTowHands(c1, c2 *Hand) bool {
	return false
}
