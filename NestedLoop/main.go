package main

import "fmt"

func main() {
  for a := 0; a < 10; a++ {
    for b := 0; b < 10; b++ {
      fmt.Println(a, "-", b)
    }
  }
}

//counts from 0 - 0 to 9 - 9
//runs inner loop (variable b) 10 times for each time it runs outer loop
//runs outter loop 10 times
