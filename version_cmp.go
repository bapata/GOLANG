package main

import (
  "fmt"
  "flag"
  "os"
  "strings"
  "strconv"
)

// [0,1] cmp [1,0] returns -1
// [1,1] cmp [1,0] returns 1
// [1,1,9] cmp [1,1] returns 1
// [1,1] cmp [1,1] returns 0

func min_int(x,y int) int {
  if x<y {
    return x
  } else {
    return y
  }
}


func cmp_int_arr(a,b []int) int {
  ii:=0
  len_a := len(a)
  len_b := len(b)
  minab := min_int(len_a,len_b)

  for ;ii<minab; {
    if a[ii]!=b[ii] {
      break
    } else {
      ii = ii+1
    }
  }
  if ii<minab {
    return a[ii]-b[ii]
  } else if ii==minab {
    if len_a >minab {
      return 1
    } else if (len_b > minab) {
      return -1
    } else {
      return a[minab-1]-b[minab-1]
    }
  }
  return 9999
}

func versionStrToIntArr(v string) []int {
    str_arr := strings.Split(v, ".")
    int_arr := make([]int, len(str_arr))
    for ii := range str_arr {
        int_arr[ii], _ = strconv.Atoi(str_arr[ii])
    }
    return int_arr
}

func main() {

  flag.Parse()
  args := flag.Args()

  if len(args) != 2 {
      fmt.Println("USAGE: ./this-script verstion-string1 version-string2\n")
      os.Exit(1);
  }
  v1 := args[0]
  v2 := args[1]
  p := versionStrToIntArr(v1)
  q := versionStrToIntArr(v2)
  // p-q
  res := cmp_int_arr(p,q)
  if res>0 {
    fmt.Println(v1 + " is greater than " + v2)
  } else if res<0 {
    fmt.Println(v2 + " is greater than " + v1)
  } else {
    fmt.Println("both (" + v1 + "),(" + v2 + ") are equal")
  }
}
