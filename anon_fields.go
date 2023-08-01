package main

// If fields of a struct are repeating, use just type (and no name)

import (
    "fmt"
)

type inner1 struct {
    repeating_field1 int
    repeating_field2 string
}

type outer1 struct {
    inner1
    outer_field1 float64
    outer_field2 string
}

func main() {

    myouter1 := outer1{
        outer_field1: 3.14,
        outer_field2: "pi",
    }

    fmt.Println(myouter1.repeating_field1)
    fmt.Println(myouter1.repeating_field2)
    fmt.Println(myouter1.outer_field1)
}
