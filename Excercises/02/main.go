package main

import "fmt"

var (
  large int
  small int
)

func main() {
  fmt.Println("Please enter a large number")
  fmt.Scan(&large)
  fmt.Println("Please enter a small number")
  fmt.Scan(&small)
  fmt.Println(large, "divided by", small, "is", large / small, "with remainder", large % small)
}
