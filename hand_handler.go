package texas_holdem

// 用来判断牌型
// 注意, 这里的牌型是指五张牌的组合, 如果不是五张牌的话, 可能会出现意想不到的bug
type IHandHandler interface {
	// 判断传入的牌是否是同花
	// @Return bool 是否是同花
	// @Return []int 如果是同花, 返回所有的牌的数字(从大到小排序, 下同)
	isFlush(cards []Card) (bool, []int)

	// 判断传入的牌是否是顺子
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	// @Return bool 是否是顺子
	// @Return []int 如果是顺子, 返回顺子的最大值即可
	isStraight(cards []int) (bool, []int)

	// 判断传入的牌是否是四条
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	// @Return bool 是否是四条
	// @Return []int 如果是四条, 返回四条的数字即可
	isFourOfAKind(cards []int) (bool, []int)

	// 判断传入的牌是否是葫芦
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	// @Return bool 是否是葫芦
	// @Return []int 如果是葫芦, 返回三条和对子数字即可(当然你也可以只返回三条数字)
	isFullHouse(cards []int) (bool, []int)

	// 判断传入的牌是否是三条
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	// @Return bool 是否是三条
	// @Return []int 如果是三条, 返回三条的数字即可
	isThreeOfAKind(cards []int) (bool, []int)

	// 判断传入的牌是否是两对
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	// @Return bool 是否是两对
	// @Return []int 如果是两对, 返回两对的数字以及单牌
	isTwoPair(cards []int) (bool, []int)

	// 判断传入的牌是否是一对
	// @Param cards 传入的数字数组, 同时该数组是已经按照从大到小的顺序排列好了
	isOnePair(cards []int) (bool, []int)
}

type HandHandler struct {
}

func (h HandHandler) isFlush(cards []Card) (bool, []int) {
	var color = cards[0].Color.Code
	for i := 1; i < len(cards); i++ {
		if cards[i].Color.Code != color {
			return false, []int{}
		}
	}
	return true, []int{cards[0].Number.Value, cards[1].Number.Value,
		cards[2].Number.Value, cards[3].Number.Value,
		cards[4].Number.Value}
}

func (h HandHandler) isStraight(cards []int) (bool, []int) {
	// 特殊情况 A5432
	if cards[0] == 14 && cards[1] == 5 && cards[2] == 4 && cards[3] == 3 && cards[4] == 2 {
		// 同时将5和A换个位置, 表示这个顺子的最大值是5
		cards[0], cards[4] = cards[4], cards[0]
		return true, []int{5}
	}

	// 其他情况
	for i := 1; i < len(cards); i++ {
		if cards[i] != cards[i-1]-1 {
			return false, []int{}
		}
	}

	return true, []int{cards[0]}
}

func (h HandHandler) isFourOfAKind(cards []int) (bool, []int) {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
		if count[v] == 4 {
			return true, []int{v}
		}
	}
	return false, []int{}
}

func (h HandHandler) isFullHouse(cards []int) (bool, []int) {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
	}
	hasThree := false
	var threeNum int
	hasTwo := false
	var twoNum int

	for k, c := range count {
		if c == 3 {
			hasThree = true
			threeNum = k
		} else if c == 2 {
			hasTwo = true
			twoNum = k
		}
	}
	if hasThree && hasTwo {
		return true, []int{threeNum, twoNum}
	}
	return false, []int{}
}

func (h HandHandler) isThreeOfAKind(cards []int) (bool, []int) {
	count := make(map[int]int)
	for _, v := range cards {
		count[v]++
		if count[v] == 3 {
			return true, []int{v}
		}
	}
	return false, []int{}
}

func (h HandHandler) isTwoPair(cards []int) (bool, []int) {
	count := make(map[int]int)
	pairs := 0

	var values []int

	for _, v := range cards {
		count[v]++
		if count[v] == 2 {
			pairs++
			values = append(values, v)
		}
	}

	if pairs == 2 {
		// 找到两对后, 需要找出单牌
		for i := 0; i < len(cards); i++ {
			if count[cards[i]] == 1 {
				values = append(values, cards[i])
				break
			}
		}
		return true, values
	}
	return false, []int{}
}

func (h HandHandler) isOnePair(cards []int) (bool, []int) {
	count := make(map[int]int)
	var numArr []int
	var twoNum = -1

	for _, v := range cards {
		count[v]++
		if count[v] == 2 {
			numArr = append(numArr, v)
			twoNum = v
			// 找出对子后, 需要找出其他三张单牌
			for i := 0; i < len(cards); i++ {
				if cards[i] != twoNum {
					numArr = append(numArr, cards[i])
				}
			}
			return true, numArr
		}
	}
	return false, []int{}
}
