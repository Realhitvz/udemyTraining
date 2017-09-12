package main

import "fmt"

func main() {
  data := []int{43, 56, 87, 12, 45, 57}
  n := average(data...)
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

//different then average package, here data functions as one item
// average(data...) the elipsis pulls each item out of the single list
//makes them multiple items so the func average will accept it
