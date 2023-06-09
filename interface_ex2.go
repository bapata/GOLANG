package main

import (
  "fmt"
)

type Animal interface {
  Speak()
  Move()
}

type Cat struct {
}

type Parrot struct {
}

func (c Cat)Speak() {
  fmt.Println("Mew")
}
func (c Cat)Move() {
  fmt.Println("Walk")
}
func (p Parrot)Speak() {
  fmt.Println("Caaaau")
}
func (c Parrot)Move() {
  fmt.Println("Fly")
}

func Live(a Animal) {
  a.Speak()
  a.Move()
}

func main() {
  a:=Cat{}
  Live(a)

  b:=Parrot{}
  Live(b)
}
