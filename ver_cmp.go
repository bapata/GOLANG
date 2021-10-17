package main

import (
  "fmt"
  "strings"
  "strconv"
)

func str2intarr(word string) []int {
  char_list := strings.Split(word, ".")
  int_list := make([]int,len(char_list))
  for index,val := range char_list {
    int_val, _ := strconv.Atoi(val)
    int_list[index]=int_val
  }
  return int_list

}

func ver_cmp(ver1,ver2 string) int {
  fmt.Println(ver1)
  fmt.Println(ver2)
  fmt.Println(str2intarr(ver1))
  return 0
}

func main() {
  fmt.Println(ver_cmp("1.10","1.2"))
}
