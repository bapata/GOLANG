package main

import "fmt"

// comments

/*
another comment
*/

// user data types
type Rectangle struct {
  leftX float64
  topY float64
  height float64
  width float64
}


func main() {
  fmt.Println("Hello World")
  var age int = 40
  var favNum float64 = 1.14345678
  randNum := 1
  // randNum = "myname"  # this will cause error
  fmt.Println(randNum)
  fmt.Println(age,favNum)

  var numOne = 1.000
  var num99 = 0.999

  fmt.Println(numOne - num99)

  const pi float64 = 3.14159265
  eee := 2.71234567
    fmt.Println(pi,eee)


  var myName string = "this is galaxy"
  fmt.Println(myName + " with a robot\n\n")

  var isOver40 bool = true
  fmt.Println(isOver40)

  fmt.Printf("%d\n\n",100)
  fmt.Printf("%f\n\n",100)
  fmt.Printf("%b\n\n",100)
  fmt.Printf("%x\n\n",100)


  // logical operators
  fmt.Println("true && false = ", true && false)
  fmt.Println(true && true)


  // loops
  i:=1
  for i<=10 {
    fmt.Println(i)
    i++
  }

  for j:=0;j<5;j++ {
    fmt.Println(j)
  }
  
    yourAge:=16
  if yourAge >= 16 {
    fmt.Println("You can drive")
  } else {
    fmt.Println("You cannot drive")
  }

  yourAge = 20
  switch yourAge {
    case 16: fmt.Println("Go Drive")
    case 18: fmt.Println("Go Vote")
    default: fmt.Println("Ho have fun")
  }

  var favNums2[5] float64

  favNums2[0] = 163
  favNums2[3] = 3.14159

  favNums3 := [5]float64 { 1, 2,3,4,5 }

  for i,value := range favNums3 {
    fmt.Println(i,value)
  }
  for _,value1 := range favNums3 {
    fmt.Println(value1,i)
  }

  fmt.Println("\n\nSLICE\n\n")
  numSlice := []int {1,3,4,5,7,8,}
  numSlice2 := numSlice[3:5]
  fmt.Println(numSlice2[0])
 x := subtractThem(4,1,2,3)
  fmt.Println("subtractThem output = ",x)

  // closure example
  num3 := 3
  doubleNum := func() int {
    num3 *= 2
    return num3
  }
  fmt.Println("Invoking doubleNum ...")
  fmt.Println(doubleNum())
  fmt.Println(doubleNum())

  // recursive
  fmt.Println("Factorial of 4 is ",factorial(4))

  // Defer demo
  defer printTwo()
  printOne()

  // Error handling
  fmt.Println(safeDiv(3,0))
  fmt.Println(safeDiv(3,2))
    fmt.Println("Invoking doubleNum ...")
  fmt.Println(doubleNum())
  fmt.Println(doubleNum())

  // recursive
  fmt.Println("Factorial of 4 is ",factorial(4))

  // Defer demo
  defer printTwo()
  printOne()

  // Error handling
  fmt.Println(safeDiv(3,0))
  fmt.Println(safeDiv(3,2))

  // Panic
  demoPanic()
  z := 0
  changeXValNow(&z)
  fmt.Println("x = ",z)
}

  
  func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }


func factorial(num int) int {
  if num == 0 { return 1 }
  return num * factorial(num-1)
}

func addThemUp(numbers []float64) float64 {
  sum := 0.0
  for _, val := range numbers {
    sum += val
  }
  return sum
}

func next2Values(number int) (int,int) {
  return number+1,number+2
}

func subtractThem(args ...int) int {
  finalValue := 0
  for _,value := range args {
    finalValue -= value
  }
  return finalValue
}

func demoPanic() {
  defer func() {
    fmt.Println(recover())
  }()
  panic("PANIC")
}

