package main

// sum of all float numbers specified in command line arguments

import (
	"fmt"
	"os"
	"strconv"
)

func strList2FloatList(a []string) []float64 {
	var retList []float64
	var f float64
	for _, v := range a {
		f, _ = strconv.ParseFloat(v, 64)
		retList = append(retList, f)
	}
	return retList
}

func recurse_sum(a []float64) float64 {
	if len(a) == 1 {
		return a[0]
	} else {
		return a[0] + recurse_sum(a[1:])
	}
}
func main() {
	arguments := os.Args[1:]
	fmt.Println(strList2FloatList(arguments))
	fmt.Println(recurse_sum(strList2FloatList(arguments)))
}
