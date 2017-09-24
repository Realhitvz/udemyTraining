package main

import "fmt"

func main() {
  mySlice := []string{
    "Hello!",
    "Bonjour",
    "Guten Abend",
    "Good Day",
    "Fair Tidings",
  }
  fmt.Print("1:3 ")
  fmt.Println(mySlice[1:3])
  fmt.Print("3: ")
  fmt.Println(mySlice[3:])
  fmt.Print(":4")
  fmt.Println(mySlice[:4])
}
