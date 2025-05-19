package readers

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/milesflo/jops/internal/types"
)

type Input struct {
}

func (i Input) Read() {}

var LoadStatusMap = map[string]types.Status{
	"Applied":             types.StatusApplied,
	"First Call Pending":  types.StatusFirstCallPending,
	"First Call Complete": types.StatusFirstCallComplete,
	"Ghosted":             types.StatusGhosted,
	"Listing Removed":     types.StatusListingRemoved,
	"Rejected":            types.StatusRejected,
}

// loadStatus converts a saved Status string to an enum
func loadStatus(statusStr string) types.Status {
	value, ok := LoadStatusMap[statusStr]

	// No ternary operator in go? Cool, cool....
	if ok {
		return value
	}
	return types.StatusPending
}

func parseDate(datestamp string) (time.Time, error) {
	return time.Parse(types.DatestampLayout, datestamp)
}

// stripNumber takes a string like "$123.4k" and returns "123.4"
func stripNumber(numSrt string) string {
	out := strings.TrimSpace(numSrt)
	out = strings.ReplaceAll(out, "K", "")
	out = strings.ReplaceAll(out, "$", "")
	return out
}

// parsePayrange will take a string like "$123K-234.5K" and return a slice of these 2 ints.
func parsePayrange(payString string) ([2]uint32, error) {
	bands := strings.Split(payString, "-")

	if len(bands) < 2 {
		return [2]uint32{}, errors.New("Malformed payrange: " + payString)
	}
	floorStr := stripNumber(bands[0])
	ceilStr := stripNumber(bands[1])
	// Extract
	floorFloat, err := strconv.ParseFloat(floorStr, 32)
	if err != nil {
		return [2]uint32{}, err
	}
	ceilFloat, err := strconv.ParseFloat(ceilStr, 32)
	if err != nil {
		return [2]uint32{}, err
	}
	fFl2 := uint32(math.Round(floorFloat * 10))
	fCl2 := uint32(math.Round(ceilFloat * 10))
	floor := fFl2 * 100
	ceil := fCl2 * 100

	return [2]uint32{floor, ceil}, nil
}
