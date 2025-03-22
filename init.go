package texas_holdem

func init() {
	ColorMap = getColorMap()
	NumberMap = getNumberMap()

	handHandler = HandHandler{}
}

func getColorMap() map[string]Color {
	colorMap := make(map[string]Color)
	colorMap[COLOR_SPADE] = Color{Code: 1, Description: "spade", Picture: "♠"}
	colorMap[COLOR_HEART] = Color{Code: 2, Description: "heart", Picture: "♥"}
	colorMap[COLOR_DIAMOND] = Color{Code: 3, Description: "diamond", Picture: "♦"}
	colorMap[COLOR_CLUB] = Color{Code: 4, Description: "club", Picture: "♣"}
	return colorMap
}

func getNumberMap() map[string]Number {
	numberMap := make(map[string]Number)

	numberMap[NUMBER_ACE] = Number{Code: 1, Description: "ace", Picture: "A", Value: 14}
	numberMap[NUMBER_2] = Number{Code: 2, Description: "2", Picture: "2", Value: 2}
	numberMap[NUMBER_3] = Number{Code: 3, Description: "3", Picture: "3", Value: 3}
	numberMap[NUMBER_4] = Number{Code: 4, Description: "4", Picture: "4", Value: 4}
	numberMap[NUMBER_5] = Number{Code: 5, Description: "5", Picture: "5", Value: 5}
	numberMap[NUMBER_6] = Number{Code: 6, Description: "6", Picture: "6", Value: 6}
	numberMap[NUMBER_7] = Number{Code: 7, Description: "7", Picture: "7", Value: 7}
	numberMap[NUMBER_8] = Number{Code: 8, Description: "8", Picture: "8", Value: 8}
	numberMap[NUMBER_9] = Number{Code: 9, Description: "9", Picture: "9", Value: 9}
	numberMap[NUMBER_10] = Number{Code: 10, Description: "10", Picture: "10", Value: 10}
	numberMap[NUMBER_JACK] = Number{Code: 11, Description: "jack", Picture: "J", Value: 11}
	numberMap[NUMBER_QUEEN] = Number{Code: 12, Description: "queen", Picture: "Q", Value: 12}
	numberMap[NUMBER_KING] = Number{Code: 13, Description: "king", Picture: "K", Value: 13}

	return numberMap
}
