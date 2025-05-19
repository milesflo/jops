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
	headerMap := buildHeaderMap(strings.Split("Company	Job Name	Link	Pay Range	Location	Status	Applied Date	Call 1 Date	Call 2 Date	Offer Date	Rejection Date", "\t"))
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
				Interviews:   nil,
				Status:       1,
				AppliedDate:  date,
			},
		},
	}

	for _, tc := range tcs {
		res, err := parseCSVRow(tc.lineText, headerMap)
		if err != nil {
			t.Error("Failed to parse row:", tc.lineText)
		}

		if diff := cmp.Diff(res, tc.expected); diff != "" {
			t.Errorf("Structs are not equal: %s", diff)
		}
	}
}
