package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuit_Card(t *testing.T) {
	tcs := []struct {
		description    string
		suit           Suit
		rank           Rank
		expectedResult Card
	}{
		{
			description:    "assert Clubs+Ace equals 'Ace of Clubs'",
			suit:           Clubs,
			rank:           Ace,
			expectedResult: CA,
		},
		{
			description:    "assert Spades+King equals 'King of Spades'",
			suit:           Spades,
			rank:           King,
			expectedResult: SK,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.suit.Card(tc.rank))
		})
	}
}

func TestSuit_Symbol(t *testing.T) {
	tcs := []struct {
		description    string
		suit           Suit
		expectedResult rune
	}{
		{
			description:    "assert Clubs.Symbol equals '♣'",
			suit:           Clubs,
			expectedResult: '♣',
		},
		{
			description:    "assert Diamonds.Symbol equals '♦'",
			suit:           Diamonds,
			expectedResult: '♦',
		},
		{
			description:    "assert Hearts.Symbol equals '♥'",
			suit:           Hearts,
			expectedResult: '♥',
		},
		{
			description:    "assert Spades.Symbol equals '♠'",
			suit:           Spades,
			expectedResult: '♠',
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.suit.Symbol())
		})
	}
}

func TestSuit_Name(t *testing.T) {
	tcs := []struct {
		description    string
		suit           Suit
		expectedResult string
	}{
		{
			description:    `assert Clubs.Name equals "Clubs"`,
			suit:           Clubs,
			expectedResult: "Clubs",
		},
		{
			description:    `assert Diamonds.Name equals "Diamonds"`,
			suit:           Diamonds,
			expectedResult: "Diamonds",
		},
		{
			description:    `assert Hearts.Name equals "Hearts"`,
			suit:           Hearts,
			expectedResult: "Hearts",
		},
		{
			description:    `assert Spades.Name equals "Spades"`,
			suit:           Spades,
			expectedResult: "Spades",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.suit.Name())
		})
	}
}

func TestSuit_String(t *testing.T) {
	tcs := []struct {
		description    string
		suit           Suit
		expectedResult string
	}{
		{
			description:    `assert Clubs.String equals "♣"`,
			suit:           Clubs,
			expectedResult: "♣",
		},
		{
			description:    `assert Diamonds.String equals "♦"`,
			suit:           Diamonds,
			expectedResult: "♦",
		},
		{
			description:    `assert Hearts.String equals "♥"`,
			suit:           Hearts,
			expectedResult: "♥",
		},
		{
			description:    `assert Spades.String equals "♠"`,
			suit:           Spades,
			expectedResult: "♠",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.suit.String())
		})
	}
}

func TestParseSuit(t *testing.T) {
	tcs := []struct {
		input          string
		expectedResult Suit
	}{
		{
			input:          "♣",
			expectedResult: Clubs,
		},
		{
			input:          "♦",
			expectedResult: Diamonds,
		},
		{
			input:          "♥",
			expectedResult: Hearts,
		},
		{
			input:          "♠",
			expectedResult: Spades,
		},
	}
	for _, tc := range tcs {
		t.Run("Parse suit: "+tc.input, func(t *testing.T) {
			suit, err := ParseSuit(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, suit)
		})
	}
}
