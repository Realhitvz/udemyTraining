package main

import (
  "fmt"
)

var balance float64
var growth float64
var interest float64

func main() {
  fmt.Println("How much money do you have?")
  fmt.Scan(&balance)
  fmt.Println("How much do you add per year?")
  fmt.Scan(&growth)
  fmt.Println("What is the interest rate?")
  fmt.Scan(&interest)
  Math()
}

func Math() {
  for i := 1; i <= 40; i++ {
    balance += growth
    balance = balance * interest
    fmt.Println("After ", i, "years you will have ", balance)
  }
}

//does a rough calculation of interest growth over 40 years
