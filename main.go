package main

import (
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filepath := os.Args[1]

	process := Process{}
	process.Input = CSVInput{Input{}, filepath}
	process.Output = TerminalOutput{}

	err := process.Load()
	must(err)

	process.Output.Write(process.Listings)
}
