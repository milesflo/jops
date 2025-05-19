package readers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/milesflo/jops/internal/types"
)

type CSVInput struct {
	Input
	Filepath string
}

func parseCSVRow(row []string) (types.JobListing, error) {
	// TODO: I just decided this is bunk. You should need an index
	//       argument to specify which header(s), if any, is in the file,
	//       and access those positions in the row.
	Company := types.Company{Name: row[0]}
	JobName := row[1]
	Link := row[2]
	// TODO
	Description := ""
	payinfo, err := parsePayrange(row[3])
	if err != nil {
		return types.JobListing{}, err
	}
	PaybandFloor := payinfo[0]
	PaybandCeil := payinfo[1]
	Location := row[4]
	Status := loadStatus(row[5])
	AppliedDate, err := parseDate(row[6])
	if err != nil {
		return types.JobListing{}, err
	}
	var OfferDate time.Time
	if len(row) > 9 {
		OfferDate, err = parseDate(row[9])
		if err != nil {
			return types.JobListing{}, err
		}
	}
	var RejectionDate time.Time
	if len(row) > 10 {
		RejectionDate, err = parseDate(row[10])
		if err != nil {
			return types.JobListing{}, err
		}
	}
	// TODO
	Interviews := []types.Interview{}
	return types.JobListing{
		Company:       Company,
		JobName:       JobName,
		Link:          Link,
		Description:   Description,
		PaybandFloor:  PaybandFloor,
		PaybandCeil:   PaybandCeil,
		Location:      Location,
		Status:        Status,
		Interviews:    Interviews,
		AppliedDate:   AppliedDate,
		OfferDate:     OfferDate,
		RejectionDate: RejectionDate,
	}, nil
}

func (s CSVInput) Read() ([]types.JobListing, error) {

	f, err := os.Open(s.Filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	_, err = csvReader.Read()
	if err != nil {
		return nil, err
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var listings []types.JobListing

	for _, line := range records {
		if len(line) == 1 {
			line = strings.Split(line[0], "\t")
		}
		fmt.Println(line)
		if len(line) < 3 {
			continue
		}
		joblisting, err := parseCSVRow(line)
		if err != nil {
			fmt.Println("Bad line found:", err)
		}
		listings = append(listings, joblisting)
	}

	return listings, nil
}
