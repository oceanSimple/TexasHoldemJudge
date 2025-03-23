package texas_holdem

const (
	COLOR_SPADE   = "spade"   // ♠
	COLOR_HEART   = "heart"   // ♥
	COLOR_DIAMOND = "diamond" // ♦
	COLOR_CLUB    = "club"    // ♣
)

const (
	NUMBER_ACE   = "ace"   // A
	NUMBER_2     = "2"     // 2
	NUMBER_3     = "3"     // 3
	NUMBER_4     = "4"     // 4
	NUMBER_5     = "5"     // 5
	NUMBER_6     = "6"     // 6
	NUMBER_7     = "7"     // 7
	NUMBER_8     = "8"     // 8
	NUMBER_9     = "9"     // 9
	NUMBER_10    = "10"    // 10
	NUMBER_JACK  = "jack"  // J
	NUMBER_QUEEN = "queen" // Q
	NUMBER_KING  = "king"  // K
)

// hand priority
const (
	HIGH_CARD       = iota //  高牌
	ONE_PAIR               //  一对
	TWO_PAIR               //  两对
	THREE_OF_A_KIND        //  三条
	STRAIGHT               // 	顺子
	FLUSH                  // 	同花
	FULL_HOUSE             //  葫芦
	FOUR_OF_A_KIND         //  四条
	STRAIGHT_FLUSH         //  同花顺
	ROYAL_FLUSH            //  皇家同花顺
)

// error string
const (
	ERROR_CARD_LENGTH_Empty = "cards can't be empty"
	ERROR_CARD_LENGTH_7     = "card length must be 7"
)
