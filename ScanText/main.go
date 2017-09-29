package main

import (
  "fmt"
  "bufio"
  "strings"
)

func main() {
  const input = "In a hole in the ground there lived a hobbit."
  scanner := bufio.NewScanner(strings.NewReader(input))
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}

//for scanner.Scan runs the for loop for each object scanned.
//bufio.ScanWords tells the program to seperate the input text into word-size objects
