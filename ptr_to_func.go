package main

import (
	"fmt"
)

type PTF func(int) int

func f1(x int) int {
	return x
}
func f2(x int) int {
	return x * x
}
func f3(x int) int {
	return x * x * x
}

func main() {
	m := map[string]PTF{"one": f1, "two": f2, "three": f3}
	fmt.Println(m)


        fmt.Println(m["one"](5))
}
