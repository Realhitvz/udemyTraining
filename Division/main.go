package main

import "fmt"

func main() {
  x := 27 / 5
  y := 27 % 5
  fmt.Println(x)
  fmt.Println(y)

  for n := 1; n < 50; n++ {
    if n % 2 == 1 {
      fmt.Println("Odd")
    } else {
      fmt.Println("Even")
    }
  }
}

// the slash denotes division
// using % instead gives us the remainder
// for loop tells us if each value up to 50 is odd or Even

// later: find out how to divide into decimal values
