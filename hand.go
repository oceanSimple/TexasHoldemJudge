package texas_holdem

import "sort"

// 定义手牌结构
type Hand struct {
	Cards []Card
	Type  int   // 牌型
	Ranks []int // 用于比较的牌面值
}

// 根据 Cards 将 Type 和 Ranks 计算出来
func (h *Hand) Handler() {
	// 计算type
	getMaxHandType(h)
}

func getMaxHandType(h *Hand) {
	h.Type = -1
	// 从七张牌中挑选五张牌，共有 C(7,5)=21 种组合。
	var combinations = getCombinations(h.Cards, 5)

	for _, combo := range combinations {
		var newType, newRanks = getTypeAndRanks(combo)
		if newType > h.Type || (h.Type == newType && compareRanks(newRanks, h.Ranks) > 0) {
			h.Type = newType
			h.Ranks = newRanks
		}
	}
}

// 获得 C( len(cards) , k) 的组合
// @param cards 牌组
// @param k 组合的大小
// @return 组合的结果
func getCombinations(cards []Card, k int) [][]Card {
	var result [][]Card                           // 存储所有组合的结果
	var backtrack func(start int, current []Card) // 定义回溯函数

	// 回溯函数，用于生成组合
	backtrack = func(start int, current []Card) {
		if len(current) == k { // 如果当前组合的长度等于k
			temp := make([]Card, k)       // 创建一个临时切片存储当前组合
			copy(temp, current)           // 复制当前组合到临时切片
			result = append(result, temp) // 将临时切片添加到结果中
			return
		}
		for i := start; i < len(cards); i++ { // 遍历所有卡片
			current = append(current, cards[i]) // 将当前卡片添加到当前组合
			backtrack(i+1, current)             // 递归调用回溯函数，生成下一个组合
			current = current[:len(current)-1]  // 回溯，移除最后一个添加的卡片
		}
	}

	backtrack(0, []Card{}) // 从第0个卡片开始，生成组合
	return result          // 返回所有组合的结果
}

func getTypeAndRanks(combination []Card) (int, []int) {
	// 将数字部分单独存入ranks中
	// 注意这里将A视为14
	var ranks = make([]int, len(combination))
	for i, card := range combination {
		ranks[i] = card.Number.Value
	}
	// 按照从大到小排序
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i] > ranks[j]
	})

	var isFlush = handHandler.isFlush(combination)
	var isStraight = handHandler.isStraight(ranks)
	if isFlush && isStraight && ranks[0] == 14 { // 皇家同花顺
		return ROYAL_FLUSH, ranks
	} else if isFlush && isStraight { // 同花顺
		return STRAIGHT_FLUSH, ranks
	}

	return -1, ranks
}

// 比较两个牌型的大小
func compareRanks(ranks1, ranks2 []int) int {
	for i := 0; i < len(ranks1); i++ {
		if ranks1[i] > ranks2[i] {
			return 1
		} else if ranks1[i] < ranks2[i] {
			return -1
		}
	}
	return 0
}
