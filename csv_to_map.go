package main

// Standalone script to parse csv file and convert to hashMap
// Usage: go run <this-go-file> <csv-file-with-2-fields>
// more input.csv
// key1,value1
// key2,value2
// key3,value3

// go run csv_parser_ex1.go ./input.csv
// map[key1:value1 key2:value2 key3:value3]

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func csvToMap(f string) map[string]string {

	ret_map := make(map[string]string)

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
		if err != nil {
			return ret_map
		}
		// fmt.Printf("Key: %s Value %s\n", record[0], record[1])
		ret_map[record[0]] = record[1]

	}
	return ret_map
}

func main() {
	myCsvFile := os.Args[1]
	ret_map := csvToMap(myCsvFile)
	fmt.Println(ret_map)
}
