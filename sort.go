package card

// Desc[ending] order by card value.
func Desc(a, b Card) bool {
	return a.Value() > b.Value()
}

// Asc[ending] order by card value.
func Asc(a, b Card) bool {
	return a.Value() < b.Value()
}

// RankDesc[ending] order by rank value.
func RankDesc(a, b Card) bool {
	return a.Rank().Value() > b.Rank().Value()
}

// RankAsc[ending] order by rank value.
func RankAsc(a, b Card) bool {
	return a.Rank().Value() < b.Rank().Value()
}

// SuitDesc[ending] order by suit value.
func SuitDesc(a, b Card) bool {
	return a.Suit().Value() > b.Suit().Value()
}

// SuitAsc[ending] order by suit value.
func SuitAsc(a, b Card) bool {
	return a.Suit().Value() < b.Suit().Value()
}
