package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

func main() {
  for i := 0; i < 10; i++ {
    resp, err := http.Get("http://localhost:8000")
    if err != nil {
      log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("Request %d: %s\n", i+1, string(body))
    time.Sleep(1 * time.Second)
  }
}

