package texas_holdem

type IHandHandler interface {
	// 判断传入的牌是否是同花
	isFlush(cards []Card) bool

	// 判断传入的牌是否是顺子
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isStraight(cards []int) bool

	// 判断传入的牌是否是四条
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isFourOfAKind(cards []int) bool

	// 判断传入的牌是否是葫芦
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isFullHouse(cards []int) bool

	// 判断传入的牌是否是三条
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isThreeOfAKind(cards []int) bool

	// 判断传入的牌是否是两对
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isTwoPair(cards []int) bool

	// 判断传入的牌是否是一对
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isOnePair(cards []int) bool
}

type HandHandler struct {
}

func (h HandHandler) isFlush(cards []Card) bool {
	var color = cards[0].Color.Code
	for i := 1; i < len(cards); i++ {
		if cards[i].Color.Code != color {
			return false
		}
	}
	return true
}

func (h HandHandler) isStraight(cards []int) bool {
	// 特殊情况 A5432
	if cards[0] == 14 && cards[1] == 5 && cards[2] == 4 && cards[3] == 3 && cards[4] == 2 {
		// 同时将5和A换个位置, 表示这个顺子的最大值是5
		cards[0], cards[4] = cards[4], cards[0]
		return true
	}

	// 其他情况
	for i := 1; i < len(cards); i++ {
		if cards[i] != cards[i-1]-1 {
			return false
		}
	}

	return true
}

func (h HandHandler) isFourOfAKind(cards []int) bool {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
		if count[v] == 4 {
			return true
		}
	}
	return false
}

func (h HandHandler) isFullHouse(cards []int) bool {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
	}
	hasThree := false
	hasTwo := false
	for _, c := range count {
		if c == 3 {
			hasThree = true
		} else if c == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func (h HandHandler) isThreeOfAKind(cards []int) bool {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
		if count[v] == 3 {
			return true
		}
	}
	return false
}

func (h HandHandler) isTwoPair(cards []int) bool {
	count := make(map[int]int)
	pairs := 0
	for _, v := range cards {
		count[v]++
		if count[v] == 2 {
			pairs++
		}
	}
	return pairs >= 2
}

func (h HandHandler) isOnePair(cards []int) bool {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
		if count[v] == 2 {
			return true
		}
	}
	return false
}
