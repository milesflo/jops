package main

import "os"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filepath := os.Args[1]
	f, err := os.ReadFile(filepath)
	must(err)
	print(string(f))
}
