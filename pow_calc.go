package main
// I wrote this to compute n to the power of n
// output at the end of the code to demonstrate what I wanted to achive
// 0 to the power of 0 is 1 (using limits)

import (
  "fmt"
  "os"
  "strconv"
  "math"
)

func main() {
  //fmt.Println("Value = ",os.Args[1])
  val_list := os.Args[1:]
  for _,v := range val_list {
    n,err := strconv.ParseFloat(v,64)
    if err!=nil {
      fmt.Println("Error, returning ...")
      return
    }
    val := math.Pow(n,n)
    if err==nil {
      fmt.Println(n,",",val)
    }
  }
}

// go run pow_calc.go 0.1 0.01 0.001 0.0001 0.00001 0.000001 0.000000001 0.00000000000001 0.000000000000000001
// 0.1 , 0.7943282347242815
// 0.01 , 0.954992586021436
// 0.001 , 0.9931160484209338
// 0.0001 , 0.9990793899844618
// 1e-05 , 0.9998848773724686
// 1e-06 , 0.9999861845848758
// 1e-09 , 0.9999999792767343
// 1e-14 , 0.9999999999996776
// 1e-18 , 1

