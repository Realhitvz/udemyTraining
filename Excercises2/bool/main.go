package main

import "fmt"

func main() {
  truth := (true && false) || (false && true) || !(false && false)
  fmt.Println(truth)
}
