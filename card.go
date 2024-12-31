package cards

import "math"

// Card in a standard 52-card deck of French-suited playing cards.
type Card uint8

// Parse the card string.
func Parse(card string) (Card, error) {
	if len(card) < 2 {
		return math.MaxUint8, ParseErr{"Parse", card, ErrSyntax}
	}
	rank, err := ParseRank(string(card[0]))
	if err != nil {
		return math.MaxUint8, err
	}
	suit, err := ParseSuit(card[1:])
	if err != nil {
		return math.MaxUint8, err
	}
	return rank.Card(suit), nil
}

// Rank of the card (suit independant value).
func (c Card) Rank() Rank {
	return Rank(int(c) % len(ranks))
}

// Suit of the card (rank independant value).
func (c Card) Suit() Suit {
	return Suit(int(c) / len(ranks))
}

// Symbol representation of the card.
func (c Card) Symbol() rune {
	return symbols[c]
}

// Name of the card (proper noun).
func (c Card) Name() string {
	return c.Rank().Name() + " of " + c.Suit().Name()
}

// String representation of the card.
func (c Card) String() string {
	return c.Rank().String() + c.Suit().String()
}

const (
	CA, C2, C3, C4, C5, C6, C7, C8, C9, CT, CJ, CQ, CK = Card(0), Card(1), Card(2), Card(3), Card(4), Card(5), Card(6), Card(7), Card(8), Card(9), Card(10), Card(11), Card(12)
	DA, D2, D3, D4, D5, D6, D7, D8, D9, DT, DJ, DQ, DK = Card(13), Card(14), Card(15), Card(16), Card(17), Card(18), Card(19), Card(20), Card(21), Card(22), Card(23), Card(24), Card(25)
	HA, H2, H3, H4, H5, H6, H7, H8, H9, HT, HJ, HQ, HK = Card(26), Card(27), Card(28), Card(29), Card(30), Card(31), Card(32), Card(33), Card(34), Card(35), Card(36), Card(37), Card(38)
	SA, S2, S3, S4, S5, S6, S7, S8, S9, ST, SJ, SQ, SK = Card(39), Card(40), Card(41), Card(42), Card(43), Card(44), Card(45), Card(46), Card(47), Card(48), Card(49), Card(50), Card(51)
)

var (
	cards = [...]Card{
		CA, C2, C3, C4, C5, C6, C7, C8, C9, CT, CJ, CQ, CK,
		DA, D2, D3, D4, D5, D6, D7, D8, D9, DT, DJ, DQ, DK,
		HA, H2, H3, H4, H5, H6, H7, H8, H9, HT, HJ, HQ, HK,
		SA, S2, S3, S4, S5, S6, S7, S8, S9, ST, SJ, SQ, SK,
	}
	symbols = [...]rune{
		'🃑', '🃒', '🃓', '🃔', '🃕', '🃖', '🃗', '🃘', '🃙', '🃚', '🃛', '🃝', '🃞',
		'🃁', '🃂', '🃃', '🃄', '🃅', '🃆', '🃇', '🃈', '🃉', '🃊', '🃋', '🃍', '🃎',
		'🂱', '🂲', '🂳', '🂴', '🂵', '🂶', '🂷', '🂸', '🂹', '🂺', '🂻', '🂽', '🂾',
		'🂡', '🂢', '🂣', '🂤', '🂥', '🂦', '🂧', '🂨', '🂩', '🂪', '🂫', '🂭', '🂮',
	}
)
