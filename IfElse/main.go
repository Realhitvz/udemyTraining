package main

import "fmt"

func main() {
  for x := 0; x <=100; x++ {
    if x >= 0 && x <= 10 {
      fmt.Println("Low Value:", x)
    } else if x > 10 && x <= 90 {
      fmt.Println("Mid Value:", x)
    } else {
      fmt.Println("High Value:", x)
    }
    }
}

//first entirely custom code
