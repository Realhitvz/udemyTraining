package main

import "fmt"

func main() {
  n := 0
  for {
    n++
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
    if n >= 50 {
      break
    }
  }
}

//prints odd numbers from 1 to 51
//loop increases n by 1, if it is even it restarts loop
//if it is odd, it prints n and checks if n >= 50
//if n is under 50 it restarts loop
