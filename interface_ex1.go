package main
import (
  "fmt"
)

// Code to show how interface{} can be used
// as a generic type and extract it as required

func main() {
  m:=map[string]interface{}{"one":1}
  m["two"]="TWO"
  m["three"]=map[int]int{3:3,33:33}
  m["four"]=4

  fmt.Println(m)

  // extract v2 as a string
  var v2 string = m["two"].(string)
  fmt.Println(v2)

  // extract interface{} as a map
  var v3 map[int]int = m["three"].(map[int]int)
  fmt.Println(v3)

  // extract interface{} as a interger
  var v4 int = m["four"].(int)
  fmt.Println(v4)
}
