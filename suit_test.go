package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuit_Value(t *testing.T) {
	tcs := []struct {
		description    string
		suit           Suit
		expectedResult int
	}{
		{
			description:    "assert Clubs value equals '0'",
			suit:           Clubs,
			expectedResult: 0,
		},
		{
			description:    "assert Diamonds value equals '1'",
			suit:           Diamonds,
			expectedResult: 1,
		},
		{
			description:    "assert Hearts value equals '2'",
			suit:           Hearts,
			expectedResult: 2,
		},
		{
			description:    "assert Spades value '3'",
			suit:           Spades,
			expectedResult: 3,
		},
		{
			description:    "assert cast integer values greater than the number of suits have expected value",
			suit:           Suit(4),
			expectedResult: 0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.suit.Value())
		})
	}
}

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
			description:    "assert Hearts.Symbol equals '❤'",
			suit:           Hearts,
			expectedResult: '❤',
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
			description:    `assert Hearts.String equals "❤"`,
			suit:           Hearts,
			expectedResult: "❤",
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
