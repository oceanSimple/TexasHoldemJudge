package texas_holdem

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type Deck struct {
	Deck []Card
}

// 获得一副牌(52张牌, 不包含大小王)
func NewDeck() Deck {
	var deck = Deck{
		Deck: obtainNewDeck(),
	}

	deck.Sort()

	return deck
}

// 将牌按照数字大小进行排序
func (deck *Deck) Sort() {
	SortCard(deck.Deck)
}

// 随机打乱牌的顺序
func (deck *Deck) Shuffle() {
	for i := len(deck.Deck) - 1; i > 0; i-- {
		jBig, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		j := int(jBig.Int64())
		deck.Deck[i], deck.Deck[j] = deck.Deck[j], deck.Deck[i]
	}
}

func (deck *Deck) String() string {
	var builder = strings.Builder{}

	for i := 0; i < len(deck.Deck); i++ {
		builder.WriteString(fmt.Sprintf("%-4s", deck.Deck[i].String()))
		if i == 12 || i == 25 || i == 38 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

// Obtain a full deck of cards (excluding kings)
func obtainNewDeck() []Card {
	var deck []Card
	for _, color := range ColorMap {
		for _, number := range NumberMap {
			deck = append(deck, NewCard(color, number))
		}
	}
	return deck
}
