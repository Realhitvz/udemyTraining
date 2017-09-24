package main

import "fmt"

func main() {
  records := make([][]string, 0)
  student1 := make([]string, 4)
  student1[0] = "Michael"
  student1[1] = "Boddy"
  student1[2] = "100.00"
  student1[3] = "88.00"
  records = append(records, student1)
  student2 := make([]string, 4)
  student2[0] = "The Shifty"
  student2[1] = "Shaft"
  student2[2] = "100.00"
  student2[3] = "70.00"
  records = append(records, student2)
  fmt.Print(records)
}
