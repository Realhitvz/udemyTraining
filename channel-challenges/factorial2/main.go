package main

import(
  "fmt"
)

type answer struct {
  x int
  y int
}

func main() {
  c := factorial(doMath(100))
  for n := range c {
    fmt.Println(n.x, "factorial is", n.y)
}
}

func doMath(n int) chan answer {
  out := make(chan answer)
  for i := n; i >= 0; i-- {
    out <- i
}
  return out
}


func factorial(n chan answer) chan answer {
  out := make(chan answer)
  go func() {
    var total answer
    total <- n
    total.y = 1
    for i := total.x; i > 1; i-- {
      total.y *= i
    }
    out <- *total
    close(out)
  }()
  return out
}
