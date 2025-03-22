package texas_holdem

import "sort"

func SortCard(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		// 先比较数字大小
		if cards[i].Number.Value != cards[j].Number.Value {
			return cards[i].Number.Value > cards[j].Number.Value
		}
		return cards[i].Color.Code < cards[j].Color.Code
	})
}
