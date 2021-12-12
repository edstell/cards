package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRank_Card(t *testing.T) {
	tcs := []struct {
		description    string
		rank           Rank
		suit           Suit
		expectedResult Card
	}{
		{
			description:    "assert Clubs+Ace equals 'Ace of Clubs'",
			rank:           Ace,
			suit:           Clubs,
			expectedResult: CA,
		},
		{
			description:    "assert Spades+King equals 'King of Spades'",
			rank:           King,
			suit:           Spades,
			expectedResult: SK,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.rank.Card(tc.suit))
		})
	}
}

func TestRank_Symbol(t *testing.T) {
	tcs := []struct {
		description    string
		rank           Rank
		expectedResult rune
	}{
		{
			description:    "assert Ace.Symbol equals 'A'",
			rank:           Ace,
			expectedResult: 'A',
		},
		{
			description:    "assert King.Symbol equals 'K'",
			rank:           King,
			expectedResult: 'K',
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.rank.Symbol())
		})
	}
}

func TestRank_Name(t *testing.T) {
	tcs := []struct {
		description    string
		rank           Rank
		expectedResult string
	}{
		{
			description:    `assert Ace.Name equals "Ace"`,
			rank:           Ace,
			expectedResult: "Ace",
		},
		{
			description:    `assert King.Name equals "King"`,
			rank:           King,
			expectedResult: "King",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.rank.Name())
		})
	}
}

func TestRank_String(t *testing.T) {
	tcs := []struct {
		description    string
		rank           Rank
		expectedResult string
	}{
		{
			description:    `assert Ace.String equals "A"`,
			rank:           Ace,
			expectedResult: "A",
		},
		{
			description:    `assert King.String equals "K"`,
			rank:           King,
			expectedResult: "K",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.rank.String())
		})
	}
}
