package writers

import (
	"os"
	"strconv"
	"strings"

	"github.com/milesflo/jops/internal/types"
)

type CSVOutput struct {
	Output
	filepath string
}

const CSVHeader = "Company	Job Name	Link	Pay Range	Location	Status	Applied Date	Call 1 Date	Call 2 Date	Offer Date	Rejection Date"

func (t CSVOutput) Write(listings []types.JobListing) error {
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
