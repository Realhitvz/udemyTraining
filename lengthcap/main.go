package main

import "fmt"

func main() {
  sample := make([]int, 0, 5)
  for i := 0; i <= 50; i++ {
    sample = append(sample, i)
    fmt.Println("Length: ", len(sample), "Cap: ", cap(sample), "Value: ", i)
  }
}
