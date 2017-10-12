package main

import "fmt"

func main() {
    c := make(chan int)
    n := 4
    total := 1
    go func() {
      for i := n; i > 0; i-- {
        total *= i
      }
      c <- total
      close(c)
    }()
    for n := range c {
      fmt.Println(n)
    }
  }
