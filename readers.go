package main

import (
	"encoding/csv"
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
	company := Company{Name: row[0]}
	jobname := row[1]
	link := row[2]
	payinfo, err := parsePayrange(row[3])
	if err != nil {
		return JobListing{}, err
	}

	Location := row[4]
	// Status := row[5]
	AppliedDate := parseDate(row[6])
	// Call1Date := ""
	// if len(row) > 7 {
	// 	Call1Date = row[7]
	// }
	// Call2Date := ""
	// if len(row) > 8 {
	// 	Call2Date = row[8]
	// }
	var OfferDate time.Time
	if len(row) > 9 {
		OfferDate = parseDate(row[9])
	}
	var RejectionDate time.Time
	if len(row) > 10 {
		RejectionDate = parseDate(row[10])
	}
	return JobListing{
		Company:       company,
		JobName:       jobname,
		Link:          link,
		PaybandFloor:  payinfo[0],
		PaybandCeil:   payinfo[1],
		Location:      Location,
		Status:        StatusApplied,
		AppliedDate:   AppliedDate,
		OfferDate:     OfferDate,
		RejectionDate: RejectionDate,
	}, nil
}

func (s CSVInput) Read() ([]JobListing, error) {

	f, err := os.Open(s.filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var listings []JobListing

	for _, line := range records {
		joblisting, _ := parseCSVRow(line)
		listings = append(listings, joblisting)
	}

	return listings, nil
}
