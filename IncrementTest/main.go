package main

import "fmt"

var x = 40

func increment() int {
  x++
  return x
}

func main() {
  fmt.Println(increment())
  fmt.Println(increment())
}

//the "increment" function is set to increase x by 1, then return x
//the first time the "main" function calls it, we get 41, the second time we get 42
