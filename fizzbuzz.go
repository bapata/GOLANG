package main
// print "fizz" if divisible by 3, "buzz" if divisible by 5, "fizzbuzz" if divisible by {3,5}, print number otherwise
import ( "fmt" )
func main() {
  var ii int
  for ii=1;ii<=100;ii++ {
    if (ii%3==0) && (ii%5==0)  {
      fmt.Println("FizzBuzz")
    } else if ii%3==0  {
      fmt.Println("Fizz")
    } else if ii%5==0  {
      fmt.Println("Buzz")
    } else {
      fmt.Println(ii)
    }
  }
}
