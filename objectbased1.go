package main

import "fmt"

type Myclass struct {
  x int
  y string
}

func (mc *Myclass) Init() {
  mc.x=10
  mc.y="ten"
}

func (mc *Myclass) Print() {
  fmt.Println(mc.x)
  fmt.Println(mc.y)
}

func main() {
  obj  := Myclass{}
  obj.Init()
  obj.Print()
}
