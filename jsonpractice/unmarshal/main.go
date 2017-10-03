package main

import (
  "fmt"
  "encoding/json"
)

type Person struct {
  First string
  Last string
  Age int `json:"wisdom score"`
}

var p1 Person

func out() {
  fmt.Println(p1.First)
  fmt.Println(p1.Last)
  fmt.Println(p1.Age)
}

func main() {
  out()
  bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
  json.Unmarshal(bs, &p1)
  out()
}
