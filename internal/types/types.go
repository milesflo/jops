package types

import "time"

const DatestampLayout = "01/02/06"

type Status int

const (
	StatusPending Status = iota
	StatusApplied
	StatusFirstCallPending
	StatusFirstCallComplete
	StatusGhosted
	StatusListingRemoved
	StatusRejected
)

type JobListing struct {
	Company       Company
	JobName       string
	Link          string
	Description   string
	PaybandFloor  uint32
	PaybandCeil   uint32
	Location      string
	Status        Status
	Interviews    []Interview
	AppliedDate   time.Time
	OfferDate     time.Time
	RejectionDate time.Time
}

type Interview struct {
	Date        time.Time
	Contact     string
	MeetingLink string
}

type Company struct {
	Name string
}

func (t Company) String() string {
	return t.Name
}
