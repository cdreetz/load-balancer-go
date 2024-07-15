package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

func main() {
  for {
    resp, err := http.Get("http://localhost:8000")
    if err != nil {
      log.Fatal(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("Response from load balancer: %s\n", string(body))
    time.Sleep(500 * time.Millisecond)
  }
}
