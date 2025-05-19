package readers

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/milesflo/jops/internal/types"
)

func TestParseRecord(t *testing.T) {
	date, err := time.Parse("1/2/06", "03/24/25")
	if err != nil {
		t.Fatal("Static test fixture failed to parse")
	}
	tcs := []struct {
		lineText []string
		expected types.JobListing
	}{
		{
			lineText: strings.Split("Wikrosoft	Software Engineer	https://wikrosoft.example	$300K - $320K	San Francisco	Applied	03/24/25	", "\t"),
			expected: types.JobListing{
				Company: types.Company{
					Name: "Wikrosoft",
				},
				JobName:      "Software Engineer",
				Link:         "https://wikrosoft.example",
				PaybandFloor: 300000,
				PaybandCeil:  320000,
				Location:     "San Francisco",
				Status:       1,
				AppliedDate:  date,
			},
		},
	}

	for _, tc := range tcs {
		res, err := parseCSVRow(tc.lineText)
		if err != nil {
			t.Error("Failed to parse row:", tc.lineText)
		}

		if diff := cmp.Diff(res, tc.expected); diff != "" {
			t.Errorf("Structs are not equal: %s", diff)
		}
	}
}

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
