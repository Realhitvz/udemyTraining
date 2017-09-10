package main

import "fmt"

const (
  _ = iota
  KB = 1 << (iota * 10)
  MB = 1 << (iota * 10)
  kilo = 1 << (10)
  mega = 1 << (20)
)

func crunch() {
  fmt.Printf("%b \n", KB)
  fmt.Printf("%b \n", MB)
  fmt.Println(KB)
  fmt.Println(MB)
  fmt.Println(kilo)
  fmt.Println(mega)
}
