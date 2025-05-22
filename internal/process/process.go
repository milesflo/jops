package process

import (
	"errors"
	"strings"

	"github.com/milesflo/jops/internal/readers"
	"github.com/milesflo/jops/internal/types"
	"github.com/milesflo/jops/internal/writers"
)

type Process struct {
	readers.Input
	writers.Output
	Listings []types.JobListing
}

func (p *Process) Load() error {
	output, err := p.Input.Read()
	p.Listings = output
	return err
}

func (p Process) Write() error {
	return p.Output.Write(p.Listings)
}

func (p Process) matches(query types.JobQuery, job types.JobListing) bool {
	if query.Company.Name != "" && query.Company.Name == job.Company.Name {
		return true
	}
	if query.JobName != "" && strings.Contains(job.JobName, query.JobName) {
		return true
	}
	if query.Location != "" && query.Location == job.Location {
		return true
	}
	if query.Status != 0 && query.Status == job.Status {
		return true
	}
	return false

}

func (p Process) Query(payload types.JobQuery, pageSize int) ([]types.JobListing, error) {
	results := make([]types.JobListing, 0)
	if pageSize < 1 {
		return []types.JobListing{}, errors.New("Invalid page size passed")
	}
	if pageSize > 20 {
		return []types.JobListing{}, errors.New("Max page size is 20")
	}
	for _, listing := range p.Listings {
		if p.matches(payload, listing) {
			results = append(results, listing)
			if len(results) >= pageSize {
				break
			}
		}
	}
	return results, nil
}
