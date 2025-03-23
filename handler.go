package texas_holdem

import "errors"

type IHandler interface {
	// 这里的card是玩家的 2张手牌 + 5张公共牌 = 一共七张牌
	CompareTwoCards(c1, c2 []Card) (int, error)

	// 比较多组牌, 返回最大牌型的索引
	// @return []int 最大牌型的索引
	GetMaxCard(cardArr [][]Card) ([]int, error)
}

type Handler struct {
	handHandler IHandHandler
}

func (h Handler) GetMaxCard(cardArr [][]Card) ([]int, error) {
	if len(cardArr) == 0 {
		return []int{-1}, errors.New(ERROR_CARD_LENGTH_Empty)
	} else if len(cardArr) == 1 {
		return []int{0}, nil
	} else {
		var maxIndex = []int{0}
		var maxCard []Card = cardArr[0]

		for i, cards := range cardArr {
			if i == 0 { // 跳过第一组
				continue
			}
			if len(cards) != 7 {
				return []int{-1}, errors.New(ERROR_CARD_LENGTH_7)
			}
			var flag, _ = h.CompareTwoCards(maxCard, cards)
			if flag < 0 {
				maxIndex = []int{i}
				maxCard = cards
			} else if flag == 0 {
				maxIndex = append(maxIndex, i)
			} else {
				// do nothing
			}
		}
		return maxIndex, nil
	}
}

func (h Handler) CompareTwoCards(c1, c2 []Card) (int, error) {
	// 首先判断是否是一共七张牌
	if len(c1) != 7 {
		return 0, errors.New(ERROR_CARD_LENGTH_7)
	}
	if len(c2) != 7 {
		return 0, errors.New(ERROR_CARD_LENGTH_7)
	}

	// 将 []Card 转换成 Hand
	var hand1 = h.cardsToHand(c1)
	var hand2 = h.cardsToHand(c2)

	return h.compareTowHands(hand1, hand2), nil
}

func (h Handler) cardsToHand(cards []Card) *Hand {
	return NewHand(cards)
}

func (h Handler) compareTowHands(c1, c2 *Hand) int {
	if c1.Type > c2.Type {
		return 1
	}
	if c1.Type < c2.Type {
		return -1
	}

	// 如果牌型相同，则比较牌面值
	var rankFlag = compareRanks(c1.Ranks, c2.Ranks)
	if rankFlag > 0 {
		return 1
	} else if rankFlag < 0 {
		return -1
	} else {
		return 0
	}
}
