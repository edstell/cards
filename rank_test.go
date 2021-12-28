package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestParseRank(t *testing.T) {
	tcs := []struct {
		input          string
		expectedResult Rank
	}{
		{
			input:          "A",
			expectedResult: Ace,
		},
		{
			input:          "2",
			expectedResult: Two,
		},
		{
			input:          "3",
			expectedResult: Three,
		},
		{
			input:          "4",
			expectedResult: Four,
		},
		{
			input:          "5",
			expectedResult: Five,
		},
		{
			input:          "6",
			expectedResult: Six,
		},
		{
			input:          "7",
			expectedResult: Seven,
		},
		{
			input:          "8",
			expectedResult: Eight,
		},
		{
			input:          "9",
			expectedResult: Nine,
		},
		{
			input:          "T",
			expectedResult: Ten,
		},
		{
			input:          "J",
			expectedResult: Jack,
		},
		{
			input:          "Q",
			expectedResult: Queen,
		},
		{
			input:          "K",
			expectedResult: King,
		},
	}
	for _, tc := range tcs {
		t.Run("Parse rank: "+tc.input, func(t *testing.T) {
			rank, err := ParseRank(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, rank)
		})
	}
}
