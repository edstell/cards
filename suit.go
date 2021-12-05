package card

// Suit of a Card (one of ♣, ♦, ❤, ♠).
type Suit uint8

// Value of the suit; bounded by the number of suits (4).
func (s Suit) Value() uint8 {
	return uint8(s) % uint8(len(suits))
}

// Card in the suit given its rank.
func (s Suit) Card(r Rank) Card {
	return Card(s.Value()*13 + r.Value())
}

// Symbol representation of the suit.
func (s Suit) Symbol() rune {
	return suitSymbols[s.Value()]
}

// Name of the suit (proper noun).
func (s Suit) Name() string {
	return suitNames[s.Value()]
}

// String representation of the suit.
func (s Suit) String() string {
	return string(s.Symbol())
}

const (
	Clubs    = Suit(0) // '♣' "Clubs"
	Diamonds = Suit(1) // '♦' "Diamonds"
	Hearts   = Suit(2) // '❤' "Hearts"
	Spades   = Suit(3) // '♠' "Spades"
)

var (
	suits       = [...]Suit{Clubs, Diamonds, Hearts, Spades}
	suitSymbols = [...]rune{'♣', '♦', '❤', '♠'}
	suitNames   = [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}
)
