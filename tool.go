package texas_holdem

import "sort"

func SortCard(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Color.Code != cards[j].Color.Code {
			return cards[i].Color.Code < cards[j].Color.Code
		}
		return cards[i].Number.Code < cards[j].Number.Code
	})
}
