package main

import "fmt"

func main() {
  greet := map[int]string{
    0: "hello",
    1: "bonjour",
    2: "guten tag",
  }
  fmt.Println(greet)
  if val, exists := greet[2]; exists {
    delete(greet, 2)
    fmt.Println("The value exists")
    fmt.Println("Val= ", val)
    fmt.Println("Exists= ", exists)
    } else {
      fmt.Println("The value does not exist")
    }

}

//in this the if function pulls the value from map key "2" before deleting it
//therefore the value still exists after we delete it from the map
