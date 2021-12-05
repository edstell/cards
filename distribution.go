package card

type Distribution func(i int, count int) bool

var Unequal = Distribution(func(i, count int) bool {
	return true
})

var Equal = Distribution(func(i, count int) bool {
	return (len(Cards) - i) > (len(Cards) % count)
})
