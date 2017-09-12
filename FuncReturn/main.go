package main

import "fmt"

func main() {
  fmt.Println(name("Michael", "Boddy"))
}

func name(fname, lname string) string {
  return fmt.Sprint(fname," ", lname)
}

//fmt.Sprint stands for String Print.
//Returns a string made from the arguments passed in
