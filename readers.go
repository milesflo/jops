package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Input struct {
}

func (i Input) Read() {}

type CSVInput struct {
	Input
	filepath string
}

const DatestampLayout = "01/02/06"

var loadStatusMap = map[string]Status{
	"Applied":             StatusApplied,
	"First Call Pending":  StatusFirstCallPending,
	"First Call Complete": StatusFirstCallComplete,
	"Ghosted":             StatusGhosted,
	"Listing Removed":     StatusListingRemoved,
	"Rejected":            StatusRejected,
}

// loadStatus converts a saved Status string to an enum
func loadStatus(statusStr string) Status {
	value, ok := loadStatusMap[statusStr]

	// No ternary operator in go? Cool, cool....
	if ok {
		return value
	}
	return StatusPending
}

func parseDate(datestamp string) time.Time {
	res, err := time.Parse(DatestampLayout, datestamp)
	must(err)
	return res
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

func parseCSVRow(row []string) (JobListing, error) {
	// TODO: I just decided this is bunk. You should need an index
	//       argument to specify which header(s), if any, is in the file,
	//       and access those positions in the row.

	Company := Company{Name: row[0]}
	JobName := row[1]
	Link := row[2]
	// TODO
	Description := ""
	payinfo, err := parsePayrange(row[3])
	if err != nil {
		return JobListing{}, err
	}
	PaybandFloor := payinfo[0]
	PaybandCeil := payinfo[1]
	Location := row[4]
	Status := loadStatus(row[5])
	AppliedDate := parseDate(row[6])
	var OfferDate time.Time
	if len(row) > 9 {
		OfferDate = parseDate(row[9])
	}
	var RejectionDate time.Time
	if len(row) > 10 {
		RejectionDate = parseDate(row[10])
	}
	// TODO
	Interviews := []Interview{}
	return JobListing{
		Company,
		JobName,
		Link,
		Description,
		PaybandFloor,
		PaybandCeil,
		Location,
		Status,
		Interviews,
		AppliedDate,
		OfferDate,
		RejectionDate,
	}, nil
}

func (s CSVInput) Read() ([]JobListing, error) {

	f, err := os.Open(s.filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	_, err = csvReader.Read()
	must(err)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var listings []JobListing

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
