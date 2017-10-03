package main

import "fmt"

type hit struct {
  str float64
}

type crit struct {
  str float64
}

type attack interface {
  damage() float64
}

func (x hit) damage() float64 {
  return x.str * 4
}

func (x crit) damage() float64 {
  return x.str * 6
}

func main() {
  v := hit{8}
  y := crit{8}
  fmt.Println("Your attack hit for", v.damage(), "damage.")
  fmt.Println("Critical Hit!", y.damage(), "damage.")
}
