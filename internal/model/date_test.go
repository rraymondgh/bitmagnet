package model

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNewDateFromIsoString(t *testing.T) {
	t.Parallel()

	type parseTest struct {
		inputString  string
		expectedDate Date
	}

	var parseTests = []parseTest{
		{
			inputString:  "2020-01-01",
			expectedDate: Date{Year: 2020, Month: 1, Day: 1},
		},
	}

	for _, test := range parseTests {
		t.Run(test.inputString, func(t *testing.T) {
			t.Parallel()

			actualOutput, err := NewDateRangeFromString(test.inputString)
			require.NoError(t, err)
			assert.Equal(t, test.expectedDate, actualOutput)
		})
	}
}
