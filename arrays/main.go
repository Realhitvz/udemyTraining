package main

import "fmt"

func main() {
  var x []string

  for n := 65; n < 100; n++ {
    x = append(x, string(n))
  }
  fmt.Println(x)
}
