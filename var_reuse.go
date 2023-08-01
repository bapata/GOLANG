package main

// code to demonstrate how to reuse variables (x in this case)

import (
	"fmt"
)

func f() int {
	return 1
}

func main() {

	y := 0
	if x := f(); x <= y {
		fmt.Println("less")
	}

	if x := f(); x > y {
		fmt.Println("greater")
	}

}
