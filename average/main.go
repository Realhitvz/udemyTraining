package main

import "fmt"

func main() {
  n := average(43, 56, 87, 12, 45, 57)
  fmt.Println(n)
}

func average(list ...int) int {
  fmt.Println(list)
  var total int
  for _, v := range list {
  total += v
}
return total / int(len(list))
}

//averages the input numbers
//on line 14, total must be on the left side of equation for it to work
