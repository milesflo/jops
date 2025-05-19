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

// Header names abstracted to keep them consistent
const (
	HeaderCompany       = "Company"
	HeaderJobName       = "Job Name"
	HeaderLink          = "Link"
	HeaderPay           = "Pay Range"
	HeaderLocation      = "Location"
	HeaderStatus        = "Status"
	HeaderAppliedDate   = "Applied Date"
	HeaderCall1Date     = "Call 1 Date"
	HeaderCall2Date     = "Call 2 Date"
	HeaderOfferDate     = "Offer Date"
	HeaderRejectionDate = "Rejection Date"
)

func parseCSVRow(row []string, headerMap map[string]int) (types.JobListing, error) {
	Company := types.Company{Name: row[headerMap[HeaderCompany]]}
	JobName := row[headerMap[HeaderJobName]]
	Link := row[headerMap[HeaderLink]]
	// TODO
	Description := ""
	payinfo, err := parsePayrange(row[headerMap[HeaderPay]])
	if err != nil {
		return types.JobListing{}, err
	}
	PaybandFloor := payinfo[0]
	PaybandCeil := payinfo[1]
	Location := row[headerMap[HeaderLocation]]
	Status := loadStatus(row[headerMap[HeaderStatus]])
	AppliedDate, err := parseDate(row[headerMap[HeaderAppliedDate]])
	if err != nil {
		return types.JobListing{}, err
	}
	var OfferDate time.Time
	if len(row) > 9 {
		OfferDate, err = parseDate(row[headerMap[HeaderOfferDate]])
		if err != nil {
			return types.JobListing{}, err
		}
	}
	var RejectionDate time.Time
	if len(row) > 10 {
		RejectionDate, err = parseDate(row[headerMap[HeaderRejectionDate]])
		if err != nil {
			return types.JobListing{}, err
		}
	}
	return types.JobListing{
		Company:       Company,
		JobName:       JobName,
		Link:          Link,
		Description:   Description,
		PaybandFloor:  PaybandFloor,
		PaybandCeil:   PaybandCeil,
		Location:      Location,
		Status:        Status,
		Interviews:    nil,
		AppliedDate:   AppliedDate,
		OfferDate:     OfferDate,
		RejectionDate: RejectionDate,
	}, nil
}

// buildHeaderMap takes the header row of the CSV file and returns a map with columns' index
func buildHeaderMap(headerRow []string) map[string]int {
	// Simple map for checking if string is in list of known headers
	headerMap := map[string]int{
		HeaderCompany:       0,
		HeaderJobName:       1,
		HeaderLink:          2,
		HeaderPay:           3,
		HeaderLocation:      4,
		HeaderStatus:        5,
		HeaderAppliedDate:   6,
		HeaderCall1Date:     7,
		HeaderCall2Date:     8,
		HeaderOfferDate:     9,
		HeaderRejectionDate: 10,
	}
	for i, val := range headerRow {
		_, ok := headerMap[val]
		if ok {
			headerMap[val] = i
		}
	}

	return headerMap
}

func (s CSVInput) Read() ([]types.JobListing, error) {
	f, err := os.Open(s.Filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	headerRow, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	headerMap := buildHeaderMap(headerRow)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var listings []types.JobListing

	for _, line := range records {
		if len(line) == 1 {
			line = strings.Split(line[0], "\t")
		}
		if len(line) < 3 {
			continue
		}
		joblisting, err := parseCSVRow(line, headerMap)
		if err != nil {
			fmt.Println("Bad line found:", err)
		}
		listings = append(listings, joblisting)
	}

	return listings, nil
}
