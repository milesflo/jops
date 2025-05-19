package writers

import (
	"fmt"
	"strings"

	"github.com/milesflo/jops/internal/types"
)

type TerminalTableOutput struct {
	Output
}

func (t TerminalTableOutput) Write(listings []types.JobListing) error {

	headers := []string{
		"Company",
		"Job Title",
		"Status",
	}

	cnMax := len(headers[0])
	jtMax := len(headers[1])
	stMax := len(headers[2])
	// Calculate padding sizes
	for _, listing := range listings {
		companyNameSize := len(listing.Company.Name)
		if companyNameSize > cnMax {
			cnMax = companyNameSize
		}
		jobnameSize := len(listing.JobName)
		if jobnameSize > jtMax {
			jtMax = jobnameSize
		}
		statusSize := len(StoreStatusMap[listing.GetStatus()])
		if statusSize > stMax {
			stMax = statusSize
		}
	}

	delimiter := "|"

	headers[0] = fmt.Sprintf("%*s", cnMax, headers[0])
	headers[1] = fmt.Sprintf("%-*s", jtMax, headers[1])
	headerLine := strings.Join(headers, delimiter)
	fmt.Println(headerLine)

	fmt.Println(strings.Repeat("-", len(headerLine)))
	for _, listing := range listings {
		company := fmt.Sprintf("%*s", cnMax, listing.Company.Name)
		jobname := fmt.Sprintf("%-*s", jtMax, listing.JobName)
		status := storeStatus(listing.GetStatus())

		fmt.Println(strings.Join([]string{company, jobname, status}, delimiter))
	}
	return nil
}

type TerminalCardOutput struct {
	Output
}

func (t TerminalCardOutput) Write(listings []types.JobListing) error {
	for _, listing := range listings {
		fmt.Printf("Company: %s\nJob Title: %s\nStatus: %s\n---\n\n", listing.Company.Name, listing.JobName, storeStatus(listing.GetStatus()))
	}
	return nil
}
