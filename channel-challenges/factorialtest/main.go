package main

import(
  "fmt"
)

func main() {
  c := factorial(doMath(10))
  for n := range c {
    fmt.Println(n)
}
}

//puts 10 * input value numbers to calculate into a channel
func doMath(n int) chan int {
  out := make(chan int)
  go func() {
  for i := 0; i < 10; i++ {
  for i := n; i >= 1; i-- {
    out <- i
}
}
  close(out)
}()
  return out
}

//takes incoming channel and finds the factorial of every number, outputs results
func factorial(n chan int) chan int {
  out := make(chan int)
  go func() {
  for x := range n {
    total := 1
    for i := x; i > 1; i-- {
      total *= i
    }
    out <- total
}
  close(out)
}()
  return out
}
