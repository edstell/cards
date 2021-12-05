package card

import "sort"

// Deck of 52 cards.
type Deck [52]Card

// NewOrderedDeck alphabetic suit and ascenting rank order.
func NewOrderedDeck() Deck {
	return Cards
}

// Sort the Deck using the provided less function. Sorting is stable; ordering
// of equal values is maintained.
func (d Deck) Sort(less func(a, b Card) bool) Deck {
	sort.SliceStable(d[:], func(i, j int) bool {
		return less(d[i], d[j])
	})
	return d
}
