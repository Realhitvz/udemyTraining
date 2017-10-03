package main

import "fmt"

type person struct {
  first string
  last string
  age   int
}

type DoubleZero struct {
  person
  LicenceToKill bool
}

func (p person) greet() {
  fmt.Println("I'm a regular person")
}

func (dz DoubleZero) greet() {
  fmt.Println("The names Bond")
}

func main() {
  p1 := DoubleZero{
    person: person{
      first: "James",
      last: "Bond",
      age: 20,
    },
    LicenceToKill: true,
  }

  p2 := person{
      first: "Miss",
      last: "Moneypenny",
      age: 18,
    }

  p1.greet()
  p2.greet()
  p1.person.greet()
}
