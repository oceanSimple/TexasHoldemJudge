package texas_holdem

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	var str = "♠A"
	runes := []rune(str)
	fmt.Println(string(runes[0])) // Output: ♠
	fmt.Println(string(runes[1])) // Output: A
}

func TestGetACard(t *testing.T) {
	var card = getACard("♠A")
	fmt.Println(card)
}
