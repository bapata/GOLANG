package main
import (
  "fmt"
)

type type1 struct {
  header int
}

type type2 struct {
  header int
  name string
}

type type3 struct {
  header int
  value float64
}

type type4 []string

type type5 map[string]string



func main() {
  m:=map[string]interface{}{
      "type1":type1{10},
      "type2":type2{20,"twenty"},
      "type3":type3{30,3.14},
      "type4":type4{"a","b"},
      "type5":type5{"ondu":"one", "yerdu":"two"},

  }
  // Walk and print each heterogenous element of the list
  for k,v := range m {
    fmt.Println("key:",k)
    fmt.Println("value:",v)
  }
  //fmt.Println(m)
}
