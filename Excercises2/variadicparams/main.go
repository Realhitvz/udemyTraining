package main

import "fmt"

func main() {
  fmt.Println(greatest(1, 2, 3, 4, 5))
}

func greatest(a ...int) int {
  var answer int
  for _, num := range a {
    if num > answer {
      answer = num
    }
  }
  return answer
}
