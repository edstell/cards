package card

import "math/rand"

// FisherYates shuffle algorithm, see
// https://en.wikipedia.org/wiki/Fisher–Yates_shuffle
func FisherYates(r rand.Rand) func(Deck) Deck {
	return func(d Deck) Deck {
		for i := len(d) - 1; i > 0; i-- {
			j := r.Intn(i + 1)
			d[i], d[j] = d[j], d[i]
		}
		return d
	}
}
