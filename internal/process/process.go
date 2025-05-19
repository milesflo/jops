package process

import (
	"github.com/milesflo/jops/internal/types"
)

type Process struct {
	Listings []types.JobListing
	Input    interface {
		Read() ([]types.JobListing, error)
	}
	Output interface {
		Write([]types.JobListing) error
	}
}

func (p *Process) Load() error {
	output, err := p.Input.Read()
	p.Listings = output
	return err
}

func (p Process) Write() error {
	return p.Output.Write(p.Listings)
}
