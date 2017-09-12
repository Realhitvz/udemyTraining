package main

import "fmt"

func main() {
  greet := func() {
    fmt.Println("Hey there!")
  }
  greet()
}

//by declaring a nameless function as a variable I can put it inside another function
//I call it with the variable
