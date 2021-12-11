package card

import "sort"

// Deck of 52 cards.
type Deck [52]Card

// OrderedDeck in alphabetic-suit and ascenting-rank order.
var OrderedDeck = Deck(cards)

// Sort the Deck using the provided less function. Sorting is stable; ordering
// of equal values is maintained.
func (d Deck) Sort(less func(a, b Card) bool) Deck {
	sort.SliceStable(d[:], func(i, j int) bool {
		return less(d[i], d[j])
	})
	return d
}

// Deal the deck to the number of hands.
func (d Deck) Deal(hands int, stop func(i int) bool) []Hand {
	out := make([]Hand, hands)
	for i := 0; i < len(d) && !stop(i); i++ {
		out[i%hands] = append(out[i%hands], d[i])
	}
	return out
}
