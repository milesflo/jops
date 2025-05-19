package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Output struct{}

type TerminalOutput struct {
	Output
}

func (t TerminalOutput) Write(listings []JobListing) error {
	for _, listing := range listings {
		fmt.Println(listing)
	}
	return nil
}

type CSVOutput struct {
	Output
	filepath string
}

const CSVHeader = "Company	Job Name	Link	Pay Range	Location	Status	Applied Date	Call 1 Date	Call 2 Date	Offer Date	Rejection Date"

func (t CSVOutput) Write(listings []JobListing) error {
	f, err := os.Create(t.filepath)
	if err != nil {
		return err
	}

	f.WriteString(CSVHeader + "\n")
	for _, l := range listings {
		payband := strconv.FormatUint(uint64(l.PaybandFloor)/1000, 10) + "K -" + strconv.FormatUint(uint64(l.PaybandCeil)/1000, 10) + "K"
		row := strings.Join([]string{l.Company.Name, l.JobName, l.Link, payband, l.Location}, ",")
		f.WriteString(row + "\n")
	}
	f.Close()
	return nil
}

var storeStatusMap = map[Status]string{
	StatusApplied:           "Applied",
	StatusFirstCallPending:  "First Call Pending",
	StatusFirstCallComplete: "First Call Complete",
	StatusGhosted:           "Ghosted",
	StatusListingRemoved:    "Listing Removed",
	StatusRejected:          "Rejected",
}

// storeStatus converts a Status value to a string for file storage
func storeStatus(status Status) string {
	value, ok := storeStatusMap[status]

	// No ternary operator in go? Cool, cool....
	if ok {
		return value
	}
	return ""
}
