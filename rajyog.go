package main

// code to determine if given date of birth has Rajayoga or not
// Ex output
// DOB: 01-01-2001
// X = 1
// Y = 5
// No rajYog. If 5,6,7 appears then has RajYog

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: ./rajyog \"mm-dd-yyyy\"\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func dobToStr(dob_mmddyyyy string) string {
	trimmed := strings.ReplaceAll(dob_mmddyyyy, "-", "")
	return trimmed
}

func strToDigits(dob string) []string {
	out := []string{}
	for i := 0; i < len(dob); i++ {
		char := string(dob[i])
		out = append(out, char)
	}
	return out
}

func addDigits(dob_array []string) int {
	sum := 0
	for _, val := range dob_array {
		val, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		sum += val
	}
	return sum
}

func addDigitsWrapper(dob_array []string) int {
	val := addDigits(dob_array)
	valAsStr := strconv.Itoa(val)
	if len(valAsStr) > 1 {
		val = addDigits(strToDigits(valAsStr))
	}
	return val
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	dob_mmddyyyy := args[0]
	bday := strings.Split(dob_mmddyyyy, "-")[0]
	bday_digit := addDigitsWrapper(strToDigits(bday))
	// reporting
	fmt.Println("DOB = " + dob_mmddyyyy)
	fmt.Println("X = " + strconv.Itoa(bday_digit))
	fmt.Println("Y = " + strconv.Itoa(addDigitsWrapper(strToDigits(dobToStr(dob_mmddyyyy)))))

}
