package writers

import (
	"github.com/milesflo/jops/internal/types"
)

type Output struct {
	Write ([]types.JobListing)
}

var StoreStatusMap = map[types.Status]string{
	types.StatusApplied:           "Applied",
	types.StatusFirstCallPending:  "First Call Pending",
	types.StatusFirstCallComplete: "First Call Complete",
	types.StatusGhosted:           "Ghosted",
	types.StatusListingRemoved:    "Listing Removed",
	types.StatusRejected:          "Rejected",
}

// storeStatus converts a Status value to a string for file storage
func storeStatus(status types.Status) string {
	value, ok := StoreStatusMap[status]

	// No ternary operator in go? Cool, cool....
	if ok {
		return value
	}
	return ""
}
