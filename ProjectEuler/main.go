package main

import "fmt"

func main() {
  var total int
  var newfib int
  fib := 2
  lastfib := 1
  for fib < 4000000 {
    if fib % 2 == 0 {
      total += fib
    }
    newfib = fib + lastfib
    lastfib = fib
    fib = newfib
    continue
  }
  fmt.Println("Sum:", total)
}
