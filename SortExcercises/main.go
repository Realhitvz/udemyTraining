package main

import (
  "sort"
  "fmt"
)

func main() {
s := []string{"Zeno", "John", "Al", "Jenny"}
sort.Sort(sort.StringSlice(s))
fmt.Println(s)
n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
sort.Sort(sort.IntSlice(n))
fmt.Println(n)
sort.Sort(sort.Reverse(sort.StringSlice(s)))
sort.Sort(sort.Reverse(sort.IntSlice(n)))
fmt.Println(s)
fmt.Println(n)
}
