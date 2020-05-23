package main

import (
	"fmt"
	"os"
	"strconv"
)

// GCD/HCF of 2 numbers
func gcd(m, n int) int {
	if m == n {
		return m
	} else if m > n {
		return gcd(m-n, n)
	} else {
		return gcd(m, n-m)
	}
}

func lcm(m, n int) int {
	var num int = 1

	if m > n {
		num = m
	} else {
		num = n
	}
	for {
		if num%m == 0 && num%n == 0 {
			break
		}
		num++
	}
	return num
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("\nUSAGE: ./mathfunc {func} [args]\n")
		os.Exit(1)
	}
	cmd := os.Args[1] // string
	arguments := os.Args[2:]

	mym, _ := strconv.Atoi(arguments[0])
	myn, _ := strconv.Atoi(arguments[1])

	if cmd == "gcd" {
		fmt.Printf("GCD of %d and %d is %d\n", mym, myn, gcd(mym, myn))
	} else if cmd == "lcm" {
		fmt.Printf("LCM of %d and %d is %d\n", mym, myn, lcm(mym, myn))
	}
}
