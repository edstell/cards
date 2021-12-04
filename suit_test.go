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
