package card

var Desc = func(a, b Card) bool {
	return a.Value() > b.Value()
}

var Asc = func(a, b Card) bool {
	return a.Value() < b.Value()
}
