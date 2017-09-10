package main

import "fmt"

func change(z *int) {
  *z = 42
}

func main() {
  x := 5
  change(&x)
  fmt.Println(x)
}

//by using & to refer to x's memory address instead of the value, paired with *
//we can change value of x using variable z
