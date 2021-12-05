package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			deck:        Cards,
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
			expectedResult: Cards,
		},
		{
			description: "assert deck is correctly sorted in ascending rank order",
			deck:        Cards,
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
			deck:        Cards,
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
			deck:        Cards,
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
			expectedResult: Cards,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.deck.Sort(tc.less))
		})
	}
}
