package main

// Implementation for collatz: https://en.wikipedia.org/wiki/Collatz_conjecture

import (
  "fmt"
  "os"
  "strconv"
)

func collatz(n int) {
  fmt.Println("Starting with ",n)
  for n>1 {
    if n%2==0 {
      n = n/2
      fmt.Println("EVEN, so n/2 = ",n)
    } else {
      n = 3*n+1
      fmt.Println("ODD, so 3n+1 = ",n)
    }
  }
}

## Usage go <script> <integer>
func main() {
  int_arg , err := strconv.Atoi(os.Args[1])
  if err != nil {
    return
  }
  collatz(int_arg)
}

// go run collatzc.go 5
//Starting with  5
//ODD, so 3n+1 =  16
//EVEN, so n/2 =  8
//EVEN, so n/2 =  4
//EVEN, so n/2 =  2
//EVEN, so n/2 =  1
