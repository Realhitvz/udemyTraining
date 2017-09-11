package main

import(
  "fmt"
  "reflect"
)

type Friend struct{
  Mary string
  Lauren string
  Steph string
}

func main() {
  switch "Lauren" {
  case "Mary":
      fmt.Println("Hi Mary")
  case "Steph":
      fmt.Println("Hey Steph")

    case "Lauren":
      fmt.Println("Hello Lauren")

    default :
      fmt.Println("Friends are elsewhere")
  }
  fmt.Println(reflect.TypeOf("Mary"))
}

//prints one of a number of options depending on the "switch" input
//can print everything below found option with "fallthrough"
