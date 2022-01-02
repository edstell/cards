package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	tcs := []struct {
		description    string
		hand           Hand
		card           Card
		expectedResult Hand
	}{
		{
			description:    "assert hand unchanged if card not present",
			card:           CA,
			hand:           []Card{C2, C3, C4},
			expectedResult: []Card{C2, C3, C4},
		},
		{
			description:    "assert card removed from front of hand",
			card:           C2,
			hand:           []Card{C2, C3, C4},
			expectedResult: []Card{C3, C4},
		},
		{
			description:    "assert card removed from middle of hand",
			card:           C3,
			hand:           []Card{C2, C3, C4},
			expectedResult: []Card{C2, C4},
		},
		{
			description:    "assert card removed from back of hand",
			card:           C4,
			hand:           []Card{C2, C3, C4},
			expectedResult: []Card{C2, C3},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.hand.Remove(tc.card))
		})
	}
}
