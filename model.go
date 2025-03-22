package texas_holdem

type Color struct {
	Code        int    // 1: spade, 2: heart, 3: diamond, 4: club
	Description string // spade, heart, diamond, club
	Picture     string // ♠, ♥, ♦, ♣
}

type Number struct {
	// 1: ace, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10, 11: jack, 12: queen, 13: king
	Code int

	// 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10, jack: 11, queen: 12, king: 13, ace: 14
	Value int

	// ace, 2, 3, 4, 5, 6, 7, 8, 9, 10, jack, queen, king
	Description string

	// A, 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K
	Picture string
}

type Card struct {
	Color  Color
	Number Number
}

func NewCard(color Color, number Number) Card {
	return Card{
		Color:  color,
		Number: number,
	}
}

func (c Card) String() string {
	return c.Color.Picture + c.Number.Picture
}
