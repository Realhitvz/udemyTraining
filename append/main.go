package main

import "fmt"

func main() {
  sliceOne := []int{1, 2, 3, 4}
  sliceTwo := []int{5, 6, 7}
  sliceOne = append(sliceOne, sliceTwo...)
  fmt.Println(sliceOne)
}
