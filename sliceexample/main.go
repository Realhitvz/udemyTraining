package main

import "fmt"

var mySlice ([]int)

func main() {
  for i := 0; i <=80; i++ {
    mySlice = append(mySlice, i)
    fmt.Println(mySlice[i], "Len:", len(mySlice), "Cap:", cap(mySlice))
  }
}
