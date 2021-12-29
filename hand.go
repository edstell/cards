package card

import "sort"

// Hand of cards.
type Hand []Card

// Sort the Hand using the provided less function. Sorting is stable; ordering
// of equal values is maintained.
func (h Hand) Sort(less func(a, b Card) bool) {
	sort.SliceStable(h, func(i, j int) bool {
		return less(h[i], h[j])
	})
}

// Contains at least one of the card.
func (h Hand) Contains(card Card) bool {
	for _, c := range h {
		if c == card {
			return true
		}
	}
	return false
}

// ContainsSuit at least once.
func (h Hand) ContainsSuit(suit Suit) bool {
	for _, c := range h {
		if c.Suit() == suit {
			return true
		}
	}
	return false
}

// ContainsRank at least once.
func (h Hand) ContainsRank(rank Rank) bool {
	for _, c := range h {
		if c.Rank() == rank {
			return true
		}
	}
	return false
}

// Find the index of the first instance of the provided card, or false if not
// in the hand.
func (h Hand) Find(card Card) (int, bool) {
	for i, c := range h {
		if c == card {
			return i, true
		}
	}
	return -1, false
}
