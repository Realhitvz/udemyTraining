package main

import "fmt"

func main() {
  n := HashBucket("Go", 12)
  fmt.Println(n)
}

func HashBucket(word string, buckets int) int {
  letter := int(word[0])
  hashed := letter % buckets
  return hashed
}
