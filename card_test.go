package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard_Suit(t *testing.T) {
	tcs := []struct {
		description    string
		card           Card
		expectedResult Suit
	}{
		{
			description:    "assert 'Ace of Clubs' is in 'Clubs'",
			card:           CA,
			expectedResult: Clubs,
		},
		{
			description:    "assert 'Ace of Diamonds' is in 'Diamonds'",
			card:           DA,
			expectedResult: Diamonds,
		},
		{
			description:    "assert 'Ace of Hearts' is in 'Hearts'",
			card:           HA,
			expectedResult: Hearts,
		},
		{
			description:    "assert 'Ace of Spades' is in 'Spades'",
			card:           SA,
			expectedResult: Spades,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.card.Suit())
		})
	}
}

func TestCard_Rank(t *testing.T) {
	tcs := []struct {
		description    string
		card           Card
		expectedResult Rank
	}{
		{
			description:    "assert 'Ace of Clubs' is an 'Ace'",
			card:           CA,
			expectedResult: Ace,
		},
		{
			description:    "assert 'Two of Clubs' is a 'Two'",
			card:           C2,
			expectedResult: Two,
		},
		{
			description:    "assert 'Three of Clubs' is a 'Three'",
			card:           C3,
			expectedResult: Three,
		},
		{
			description:    "assert 'Four of Clubs' is a 'Four'",
			card:           C4,
			expectedResult: Four,
		},
		{
			description:    "assert 'Five of Clubs' is a 'Five'",
			card:           C5,
			expectedResult: Five,
		},
		{
			description:    "assert 'Six of Clubs' is a 'Six'",
			card:           C6,
			expectedResult: Six,
		},
		{
			description:    "assert 'Seven of Clubs' is a 'Seven'",
			card:           C7,
			expectedResult: Seven,
		},
		{
			description:    "assert 'Eight of Clubs' is an 'Eight'",
			card:           C8,
			expectedResult: Eight,
		},
		{
			description:    "assert 'Nine of Clubs' is a 'Nine'",
			card:           C9,
			expectedResult: Nine,
		},
		{
			description:    "assert 'Ten of Clubs' is a 'Ten'",
			card:           CT,
			expectedResult: Ten,
		},
		{
			description:    "assert 'Jack of Clubs' is a 'Jack'",
			card:           CJ,
			expectedResult: Jack,
		},
		{
			description:    "assert 'Queen of Clubs' is a 'Queen'",
			card:           CQ,
			expectedResult: Queen,
		},
		{
			description:    "assert 'King of Clubs' is a 'King'",
			card:           CK,
			expectedResult: King,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.card.Rank())
		})
	}

}