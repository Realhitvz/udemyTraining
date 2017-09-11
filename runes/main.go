package main

import "fmt"

func main() {
  for i := 50; i <= 150; i++ {
    fmt.Println(i, "-", string(i), "-", []byte(string(i)))
  }
}

//for numbers 50-150 prints the number, the character represented by that number in UTF-8
//and the byte(s) to form that number
