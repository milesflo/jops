package readers

import "testing"

func TestParsePayRange(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected [2]uint32
	}{
		{
			name:     "Small test",
			input:    "$123K-$345K",
			expected: [2]uint32{123000, 345000},
		},
		{
			name:     "Space in hyphen",
			input:    "123.5K - 185.4K",
			expected: [2]uint32{123500, 185400},
		},
		{
			name:     "Big numbers..?",
			input:    "$800.5K - $1085.43K",
			expected: [2]uint32{800500, 1085400},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := parsePayrange(tc.input)
			if res != tc.expected {
				t.Errorf("got %v, expected %v", res, tc.expected)
			}
		})
	}
}

func TestStripNumber(t *testing.T) {
	tcs := []struct {
		input    string
		expected string
	}{
		{
			input:    "$123.4K",
			expected: "123.4",
		},
		{
			// Throw this in for good measure
			input:    "123.4",
			expected: "123.4",
		},
	}

	for _, tc := range tcs {
		res := stripNumber(tc.input)

		if res != tc.expected {
			t.Errorf("[stripNumber] got %v, expected %v", res, tc.expected)
		}
	}
}
