package main

// Standalone script to parse csv file and convert to hashMap
// Usage: go run <this-go-file> <csv-file-with-multiple-fields>
// more input.csv
// keyA,valueA1,valueA2
// keyB,valueB1,valueB2,valueB3
// keyC,valueC1

// go run csv_parser_ex1.go ./input.csv
// map[keyA: [valueA1] keyB:[valueB1,valueB2 ..],  keyC:[valueC1,valueC2]

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func csvToMap(f string) map[string][]string {

	ret_map := make(map[string][]string)

	mycsvfile, err := os.Open(f)
	if err != nil {
		return ret_map
	}

	// Parse the file
	r := csv.NewReader(mycsvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
                // key is one element, value is list of elements
		ret_map[record[0]] = record[1:]

	}
	return ret_map
}

func main() {
	myCsvFile := os.Args[1]
	ret_map := csvToMap(myCsvFile)
	//fmt.Println(ret_map)
        for k,v := range ret_map {
          fmt.Println("Key:",k,",Value:",v)
        }
}
