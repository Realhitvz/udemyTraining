package main

import "fmt"

func main() {
  half := func(n int) (int, bool) {
    if n%2 == 0 {
      return n/2, true
    } else {
      return n/2, false
    }
  }
  fmt.Println(half(5))
}
