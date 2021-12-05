package card

// Rank of a card (suit independant value).
type Rank uint8

// Value of the rank; bounded by the number of ranks (13).
func (r Rank) Value() uint8 {
	return uint8(r) % uint8(len(ranks))
}

// Card with corresponding rank given its suit.
func (r Rank) Card(s Suit) Card {
	return Card(s.Value()*13 + r.Value())
}

// Symbol representation of the rank.
func (r Rank) Symbol() rune {
	return rankSymbols[r.Value()]
}

// Name of the rank (proper noun).
func (r Rank) Name() string {
	return rankNames[r.Value()]
}

// String representation of the rank.
func (r Rank) String() string {
	return string(r.Symbol())
}

const (
	Ace   = Rank(0)  // 'A' "Ace"
	Two   = Rank(1)  // '2' "Two"
	Three = Rank(2)  // '3' "Three"
	Four  = Rank(3)  // '4' "Four"
	Five  = Rank(4)  // '5' "Five"
	Six   = Rank(5)  // '6' "Six"
	Seven = Rank(6)  // '7' "Seven"
	Eight = Rank(7)  // '8' "Eight"
	Nine  = Rank(8)  // '9' "Nine"
	Ten   = Rank(9)  // 'T' "Ten"
	Jack  = Rank(10) // 'J' "Jack"
	Queen = Rank(11) // 'Q' "Queen"
	King  = Rank(12) // 'K' "King"
)

var (
	ranks       = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	rankSymbols = [...]rune{'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K'}
	rankNames   = [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
)
