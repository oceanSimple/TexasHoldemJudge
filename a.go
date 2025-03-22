package texas_holdem

// 获取一张牌
// 推荐参数使用static内定义好的常量, 例如:COLOR_SPADE, COLOR_CLUB, NUMBER_ACE, NUMBER_9, NUMBER_KING
func GetACard(colorStr, numberStr string) Card {
	var color = GetColor(colorStr)
	var number = GetNumber(numberStr)

	return Card{
		Color:  color,
		Number: number,
	}
}

// 获取Color
// 推荐参数使用static内定义好的常量, 例如:COLOR_SPADE, COLOR_CLUB
func GetColor(colorStr string) Color {
	return ColorMap[colorStr]
}

// 获取Number
// 推荐参数使用static内定义好的常量, 例如:NUMBER_ACE, NUMBER_9, NUMBER_KING
func GetNumber(numberStr string) Number {
	return NumberMap[numberStr]
}

// 获取一副牌
func GetADeck() Deck {
	return NewDeck()
}

// 获得处理器
func GetAHandler() IHandler {
	return Handler{
		handHandler: HandHandler{},
	}
}
