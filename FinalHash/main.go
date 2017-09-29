package main

import (
  "fmt"
  "bufio"
  "net/http"
  "log"
)

func main() {
  res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(res.Body)
  defer res.Body.Close()
  scanner.Split(bufio.ScanWords)
  buckets := make([][]string, 12)
  //creates slice which holds slices
  for i := 0; i < 12; i++ {
    buckets = append(buckets, []string{})
  }
  //creates the slices to put in the above slice
  for scanner.Scan() {
    word := scanner.Text()
    n := HashBucket(word, 12)
    buckets[n] = append(buckets[n], word)
  }
  //seperates each word of source text into the sub-slices
  for i := 0; i < 12; i++ {
    fmt.Println(i, " - ", len(buckets[i]))
  }
  fmt.Println(buckets[6])
}

func HashBucket(word string, tbuck int) int {
  var sum int
  for _, v := range word {
    sum += int(v)
  }
  return sum % tbuck
}
