package main

import (
	"encoding/csv"
	"os"
)

func main() {
	data := [][]string{
		{"first_name", "last_name", "username"},
		{"Keiji", "Miyazaki", "miyahkun"},
	}

	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(f)

	for _, record := range data {
		if err := w.Write(record); err != nil {
			panic(err)
		}
	}
	w.Flush()
}
