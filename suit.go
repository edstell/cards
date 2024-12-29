package cards

// Suit of a Card (one of ♣, ♦, ♥, ♠).
type Suit int

// ParseSuit string to a Suit.
func ParseSuit(suit string) (Suit, error) {
	switch suit {
	case "♣":
		return Clubs, nil
	case "♦":
		return Diamonds, nil
	case "♥":
		return Hearts, nil
	case "♠":
		return Spades, nil
	default:
		return -1, ParseErr{"ParseSuit", suit, ErrSyntax}
	}
}

// Card in the suit given its rank.
func (s Suit) Card(r Rank) Card {
	return Card(int(s)*13 + int(r))
}

// Symbol representation of the suit.
func (s Suit) Symbol() rune {
	return suitSymbols[s]
}

// Name of the suit (proper noun).
func (s Suit) Name() string {
	return suitNames[s]
}

// String representation of the suit.
func (s Suit) String() string {
	return string(s.Symbol())
}

const (
	Clubs    = Suit(0) // '♣' "Clubs"
	Diamonds = Suit(1) // '♦' "Diamonds"
	Hearts   = Suit(2) // '♥' "Hearts"
	Spades   = Suit(3) // '♠' "Spades"
)

var (
	suits       = [...]Suit{Clubs, Diamonds, Hearts, Spades}
	suitSymbols = [...]rune{'♣', '♦', '♥', '♠'}
	suitNames   = [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}
)
