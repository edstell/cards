package card

// All cards should be dealt; don't stop.
func All(int) bool {
	return false
}

// Equal hands; stop when no more equal hands can be dealt.
func Equal(hands int) func(int) bool {
	return func(i int) bool {
		return (len(OrderedDeck) % hands) >= (len(OrderedDeck) - i)
	}
}

// Num[ber] of cards to deal in total; stop when 'num' cards has been reached.
func Num(num int) func(int) bool {
	return func(i int) bool {
		return num == i
	}
}

// Size of hands to deal; stop when each hand contains 'size' cards.
func Size(hands, size int) func(int) bool {
	return func(i int) bool {
		return i >= size*hands
	}
}
