package main

import (
    "fmt"
)

type StructParent interface {
  PrintFields()
}

type MyStruct1 struct {
    Myfield int
}

func (s MyStruct1) PrintFields() {
    fmt.Println("I am in PrintFields of MyStruct1")
    fmt.Println("Field1:", s.Myfield)
}

type MyStruct2 struct {
    Myfield string
}
func (s MyStruct2) PrintFields() {
    fmt.Println("I am in PrintFields of MyStruct2")
    fmt.Println("Field1:", s.Myfield)
}

// printStructure prints the fields of any structure
func printStructure(s StructParent) {
    // Use type switch to handle different types
    s.PrintFields()
}


func main() {

    // Example with a custom struct
    s1 := MyStruct1{Myfield: 42}
    s2 := MyStruct2{Myfield: "world"}
    printStructure(s1)
    printStructure(s2)
}

