package main

import "time"

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

type Process struct {
	Listings []JobListing
	Input    interface {
		Read() ([]JobListing, error)
	}
	Output interface {
		Write([]JobListing) error
	}
}

func (p Process) Load() error {
	output, err := p.Input.Read()
	p.Listings = output
	return err
}

func (p Process) Write() error {
	return p.Output.Write(p.Listings)
}
