package main

import "fmt"

func main() {
  greet := map[int]string{
    0: "hello",
    1: "bonjour",
    2: "guten tag",
  }
  for key, val := range greet {
    fmt.Println(key, " - ", val)
  }
}
