package process

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/milesflo/jops/internal/types"
)

func TestQuery(t *testing.T) {
	tcs := []struct {
		name     string
		listings []types.JobListing
		query    types.JobListing
		expected []types.JobListing
		pageSize int
	}{
		{
			name: "Company name query",
			listings: []types.JobListing{
				{
					Company: types.Company{
						Name: "Foogle",
					},
				},
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
				},
				{
					Company: types.Company{
						Name: "Bookface",
					},
				},
			},
			query: types.JobListing{
				Company: types.Company{
					Name: "Wikiwho",
				},
			},
			expected: []types.JobListing{
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
				},
			},
			pageSize: 10,
		},

		{
			name: "pageSize clip",
			listings: []types.JobListing{
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
					JobName: "Janitor",
				},
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
					JobName: "CTO",
				},
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
					JobName: "Software Engineer",
				},
			},
			query: types.JobListing{
				Company: types.Company{
					Name: "Wikiwho",
				},
			},
			expected: []types.JobListing{
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
					JobName: "Janitor",
				},
				{
					Company: types.Company{
						Name: "Wikiwho",
					},
					JobName: "CTO",
				},
			},
			pageSize: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := Process{
				Listings: tc.listings,
			}

			res, err := p.Query(tc.query, tc.pageSize)

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(res, tc.expected); diff != "" {
				t.Errorf("Structs are not equal: %s", diff)
			}
		})
	}
}
