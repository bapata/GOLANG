package main

// script to reverse a statement
// go run <this-go-code>

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "this is the statement to reverse"
	c := []string{}
	b := strings.Fields(sentence)
	for _, w := range b {
		c = append([]string{w}, c...)
	}
	fmt.Println(c)

}
