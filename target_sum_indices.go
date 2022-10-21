package main

import (
  "fmt"
)

type MyTargetLists struct {
  input_list []int
  target int
  target_index_list []int
}

func (tl *MyTargetLists )set_target(t int) {
  tl.target = t
}

func (tl *MyTargetLists )set_input_list(mylist []int) {
  tl.input_list= mylist
}

// algorithm
func (tl *MyTargetLists )get_index_list() []int {
  return tl.target_index_list
}

func main() {
  mtl := MyTargetLists{}
  mtl.set_target(10)
  mtl.set_input_list([]int{10,20})
  fmt.Println(mtl)
}
