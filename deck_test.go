package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeck_Sort(t *testing.T) {
	tcs := []struct {
		description    string
		deck           Deck
		less           func(a, b Card) bool
		expectedResult Deck
	}{
		{
			description: "assert deck is correctly sorted in descending order",
			deck:        cards,
			less:        Desc,
			expectedResult: Deck{
				SK, SQ, SJ, ST, S9, S8, S7, S6, S5, S4, S3, S2, SA,
				HK, HQ, HJ, HT, H9, H8, H7, H6, H5, H4, H3, H2, HA,
				DK, DQ, DJ, DT, D9, D8, D7, D6, D5, D4, D3, D2, DA,
				CK, CQ, CJ, CT, C9, C8, C7, C6, C5, C4, C3, C2, CA,
			},
		},
		{
			description: "assert deck is correctly sorted in ascending order",
			deck: Deck{
				SK, SQ, SJ, ST, S9, S8, S7, S6, S5, S4, S3, S2, SA,
				HK, HQ, HJ, HT, H9, H8, H7, H6, H5, H4, H3, H2, HA,
				DK, DQ, DJ, DT, D9, D8, D7, D6, D5, D4, D3, D2, DA,
				CK, CQ, CJ, CT, C9, C8, C7, C6, C5, C4, C3, C2, CA,
			},
			less:           Asc,
			expectedResult: cards,
		},
		{
			description: "assert deck is correctly sorted in ascending rank order",
			deck:        cards,
			less:        RankAsc,
			expectedResult: Deck{
				CA, DA, HA, SA,
				C2, D2, H2, S2,
				C3, D3, H3, S3,
				C4, D4, H4, S4,
				C5, D5, H5, S5,
				C6, D6, H6, S6,
				C7, D7, H7, S7,
				C8, D8, H8, S8,
				C9, D9, H9, S9,
				CT, DT, HT, ST,
				CJ, DJ, HJ, SJ,
				CQ, DQ, HQ, SQ,
				CK, DK, HK, SK,
			},
		},
		{
			description: "assert deck is correctly sorted in descending rank order",
			deck:        cards,
			less:        RankDesc,
			expectedResult: Deck{
				CK, DK, HK, SK,
				CQ, DQ, HQ, SQ,
				CJ, DJ, HJ, SJ,
				CT, DT, HT, ST,
				C9, D9, H9, S9,
				C8, D8, H8, S8,
				C7, D7, H7, S7,
				C6, D6, H6, S6,
				C5, D5, H5, S5,
				C4, D4, H4, S4,
				C3, D3, H3, S3,
				C2, D2, H2, S2,
				CA, DA, HA, SA,
			},
		},
		{
			description: "assert deck is correctly sorted in descending suit order",
			deck:        cards,
			less:        SuitDesc,
			expectedResult: Deck{
				SA, S2, S3, S4, S5, S6, S7, S8, S9, ST, SJ, SQ, SK,
				HA, H2, H3, H4, H5, H6, H7, H8, H9, HT, HJ, HQ, HK,
				DA, D2, D3, D4, D5, D6, D7, D8, D9, DT, DJ, DQ, DK,
				CA, C2, C3, C4, C5, C6, C7, C8, C9, CT, CJ, CQ, CK,
			},
		},
		{
			description: "assert deck is correctly sorted in ascending suit order",
			deck: Deck{
				SA, S2, S3, S4, S5, S6, S7, S8, S9, ST, SJ, SQ, SK,
				HA, H2, H3, H4, H5, H6, H7, H8, H9, HT, HJ, HQ, HK,
				DA, D2, D3, D4, D5, D6, D7, D8, D9, DT, DJ, DQ, DK,
				CA, C2, C3, C4, C5, C6, C7, C8, C9, CT, CJ, CQ, CK,
			},
			less:           SuitAsc,
			expectedResult: cards,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.deck.Sort(tc.less))
		})
	}
}

func TestDeck_Deal(t *testing.T) {
	tcs := []struct {
		description string
		deck        Deck
		hands       int
		stop        func(int) bool
		assert      func(*testing.T, []Hand)
	}{
		{
			description: "assert equal hands are dealt",
			deck:        OrderedDeck(),
			hands:       5,
			stop:        Equal(5),
			assert: func(t *testing.T, hands []Hand) {
				require.Len(t, hands, 5)
				for _, hand := range hands {
					assert.Len(t, hand, 10)
				}
			},
		},
		{
			description: "assert requested number of cards are dealt",
			deck:        OrderedDeck(),
			hands:       1,
			stop:        Num(20),
			assert: func(t *testing.T, hands []Hand) {
				require.Len(t, hands, 1)
				assert.Len(t, hands[0], 20)
			},
		},
		{
			description: "assert hands are of size requested",
			deck:        OrderedDeck(),
			hands:       3,
			stop:        Size(3, 17),
			assert: func(t *testing.T, hands []Hand) {
				require.Len(t, hands, 3)
				for _, hand := range hands {
					assert.Len(t, hand, 17)
				}
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			tc.assert(t, tc.deck.Deal(tc.hands, tc.stop))
		})
	}
}
