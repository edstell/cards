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
