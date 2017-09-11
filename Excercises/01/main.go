package main

import "fmt"

var name string

func main() {
  fmt.Println("What is your name?")
  fmt.Scan(&name)
  fmt.Println("Hello", name)
}
