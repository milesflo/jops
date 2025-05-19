package process

import (
	"errors"

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

func (p Process) Query(payload types.JobListing, pageSize int) ([]types.JobListing, error) {
	results := make([]types.JobListing, 0)
	if pageSize < 1 {
		return []types.JobListing{}, errors.New("Invalid page size passed")
	}
	if pageSize > 20 {
		return []types.JobListing{}, errors.New("Max page size is 20")
	}
	for _, listing := range p.Listings {
		if payload.Company.Name == listing.Company.Name {
			results = append(results, listing)
			if len(results) >= pageSize {
				break
			}
		}
	}
	return results, nil
}
