package main

import "fmt"


func main() {
  total := 0
  for x := 1; x <= 999; x++ {
    if x % 3 == 0 {
      total = total + x
    } else if x % 5 == 0 {
      total = total + x
    } else {
      continue
    }
  }
  fmt.Println(total)
}
