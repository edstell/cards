package card

// Distribution operator should return false when the desired deal distribution
// is met.
type Distribution func(i int, count int) bool

// Unequal distribution; deal all cards.
var Unequal = Distribution(func(i, count int) bool {
	return true
})

// Equal distribution; stop dealing when largest count multiple is reached.
var Equal = Distribution(func(i, count int) bool {
	return (len(Cards) - i) > (len(Cards) % count)
})
