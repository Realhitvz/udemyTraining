package main

import (
  "bufio"
  "fmt"
  "log"
  "net/http"
)

var hashbuckets int = 12

func main() {
  res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(res.Body)
  defer res.Body.Close()
  scanner.Split(bufio.ScanWords)
  buckets := make([]int, hashbuckets)
  for scanner.Scan() {
    n := HashBucket(scanner.Text())
    buckets[n]++
  }
  fmt.Println(buckets)
}

func HashBucket(word string) int {
  var sum int
  for _, v := range word {
    sum += int (v)
  }
  return sum % hashbuckets
}
