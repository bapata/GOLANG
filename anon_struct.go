package main

// code to demo use of anon struct

import (
	"fmt"
)

func main() {
	mys := struct {
		a int
		b string
	}{a: 10, b: "c"}

	myslist := []struct {
		c float32
		d int32
	}{{3.14, 44}, {2.78, 5}}

	fmt.Println(mys)
	fmt.Println(myslist)
}
