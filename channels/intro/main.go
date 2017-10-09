package main

import (
  "fmt"
  "time"
)

func main() {
  c := make(chan int)

  go func(){
    for i := 0; i < 10; i++ {
      c <- i
    }
    close(c)
  }()

  go func(){
    for n := range c {
      fmt.Println(n)
    }
  }()
  time.Sleep(time.Second)
}

//the println cannot take the value from the channel until it has been passed in by previous function
