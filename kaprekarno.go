package main

// Kaprekar constant - 6174
// USAGE: go run <this-code> <4-digit-number>

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "sort"
)

func high_to_low(n int) int {
  n_as_str := fmt.Sprintf("%d", n)
  n_as_str_list := strings.Split(n_as_str,"")
  sort.Strings(n_as_str_list)
  sort.Sort(sort.Reverse(sort.StringSlice(n_as_str_list)))
  ret_str := strings.Join(n_as_str_list, "")
  n,err := strconv.Atoi(ret_str)
  if err!=nil {
    fmt.Println("error l2h")
    return -1
  }
  return n
}

func low_to_high(n int) int {
  n_as_str := fmt.Sprintf("%d", n)
  n_as_str_list := strings.Split(n_as_str,"")
  sort.Strings(n_as_str_list)
  ret_str := strings.Join(n_as_str_list, "")
  n,err := strconv.Atoi(ret_str)
  if err!=nil {
    return -1
  }
  return n
}

func main() {
  args := os.Args[1:]
  if len(args)>1 {
    os.Exit(1)
  }
  n, err := strconv.Atoi(args[0])
  if err != nil {
    return
  }
  if (n<1000) || (n>9999) {
    fmt.Println("Digit has to be between 1000 and 9999 ..")
    os.Exit(1)
  }

  knum_constant := 6174
  fmt.Println("You started with initial number = ",n)
  a:=0
  b:=0

  for n != knum_constant {
    a=high_to_low(n)
    fmt.Println("high_to_low: ",a)
    b=low_to_high(n)
    fmt.Println("low_to_high: ",b)
    n=a-b
    if(n<=0) {
      fmt.Println("digits should be unique, please try again")
      os.Exit(1)
    }
    fmt.Println("intermediate number is: ",n)
  }
  fmt.Println("We have hit kaprekars number: ",n)
}
