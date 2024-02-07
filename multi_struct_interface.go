package main

import (
    "fmt"
)

type StructParent interface {
  PrintFields()
}

type MyStruct1 struct {
    Field1 int
}

func (s MyStruct1) PrintFields() {
    fmt.Println("I am in PrintFields of MyStruct1")
    fmt.Println("Field1:", s.Field1)
}

type MyStruct2 struct {
    Field2 string
}
func (s MyStruct2) PrintFields() {
    fmt.Println("I am in PrintFields of MyStruct2")
    fmt.Println("Field2:", s.Field2)
}

// printStructure prints the fields of any structure
func printStructure(s StructParent) {
    // Use type switch to handle different types
    s.PrintFields()
}


func main() {

    // Example with a custom struct
    s1 := MyStruct1{Field1: 42}
    s2 := MyStruct2{Field2: "world"}
    printStructure(s1)
    printStructure(s2)
}
