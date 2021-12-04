package card

import "sort"

type Deck [52]Card

func NewOrderedDeck() Deck {
	return Cards
}

// Sort the Deck using the provided less function.
func (d Deck) Sort(less func(a, b Card) bool) Deck {
	sort.Slice(d[:], func(i, j int) bool {
		return less(d[i], d[j])
	})
	return d
}
