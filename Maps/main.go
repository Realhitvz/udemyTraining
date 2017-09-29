package main

import "fmt"

func main() {
  var greeting = make(map[string]string)
  greeting["Mike"] = "Good Morning"
  greeting["Lyssa"] = "Good Afternoon"
  fmt.Println(greeting)
  greeting["Other"] = "Good Night"
  fmt.Println(greeting)
  delete(greeting, "Other")
  fmt.Println(greeting)
}
