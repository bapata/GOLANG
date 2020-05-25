package main
// min,max from command line arguments
// USAGE: <this-exe> 1 2 -3 -1
// OR
// go run <this-go-code> 1 2 -3 -1

import (
	"fmt"
	"os"
	"strconv"
)

// min,max:=min_max(a) where a=["3.14","1.1","2.0"..]
func min_max(args []string) (float64, float64) {
	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)
	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("One or more args required")
		os.Exit(1)
	}

	args := os.Args
	min, max := min_max(args)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
