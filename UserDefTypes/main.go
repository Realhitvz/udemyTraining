package main

import "fmt"

type foo int

func main() {
  var myAge foo
  myAge = 26
  fmt.Printf("%T %v \n", myAge, myAge)
  var yourAge int
  yourAge = 44
  fmt.Println(int(myAge) + yourAge)
}
