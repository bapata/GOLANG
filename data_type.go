package main

import (
  "fmt"
  "reflect"
)

func main() {
  a := make(map[string]int)
  b := "abracadabra"
  c := 10.2
  d := 11
  e := 3.14159
  fmt.Println(reflect.TypeOf(a))
  fmt.Println(reflect.TypeOf(b))
  fmt.Println(reflect.TypeOf(c))
  fmt.Println(reflect.TypeOf(d))
  fmt.Println(reflect.TypeOf(e))
}
