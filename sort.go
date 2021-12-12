package card

// Desc[ending] order by card value.
func Desc(a, b Card) bool {
	return a > b
}

// Asc[ending] order by card value.
func Asc(a, b Card) bool {
	return a < b
}

// RankDesc[ending] order by rank value.
func RankDesc(a, b Card) bool {
	return a.Rank() > b.Rank()
}

// RankAsc[ending] order by rank value.
func RankAsc(a, b Card) bool {
	return a.Rank() < b.Rank()
}

// SuitDesc[ending] order by suit value.
func SuitDesc(a, b Card) bool {
	return a.Suit() > b.Suit()
}

// SuitAsc[ending] order by suit value.
func SuitAsc(a, b Card) bool {
	return a.Suit() < b.Suit()
}
