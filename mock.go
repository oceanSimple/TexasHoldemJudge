package texas_holdem

func getACard(str string) *Card {
	runes := []rune(str)
	var color = getColor(string(runes[0]))   // Output: ♠
	var number = getNumber(string(runes[1])) // Output: A

	return &Card{
		Color:  color,
		Number: number,
	}
}

func getColor(color string) Color {
	switch color {
	case "♠":
		return ColorMap[COLOR_SPADE]
	case "♥":
		return ColorMap[COLOR_HEART]
	case "♦":
		return ColorMap[COLOR_DIAMOND]
	case "♣":
		return ColorMap[COLOR_CLUB]
	}
	return ColorMap[COLOR_SPADE]
}

func getNumber(number string) Number {
	switch number {
	case "A":
		return NumberMap[NUMBER_ACE]
	case "K":
		return NumberMap[NUMBER_KING]
	case "Q":
		return NumberMap[NUMBER_QUEEN]
	case "J":
		return NumberMap[NUMBER_JACK]
	case "10":
		return NumberMap[NUMBER_10]
	case "9":
		return NumberMap[NUMBER_9]
	case "8":
		return NumberMap[NUMBER_8]
	case "7":
		return NumberMap[NUMBER_7]
	case "6":
		return NumberMap[NUMBER_6]
	case "5":
		return NumberMap[NUMBER_5]
	case "4":
		return NumberMap[NUMBER_4]
	case "3":
		return NumberMap[NUMBER_3]
	case "2":
		return NumberMap[NUMBER_2]
	default:
		return NumberMap[number]
	}
}
