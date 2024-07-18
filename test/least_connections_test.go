package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "sync"
  "time"
)

func startServer(port string, processingTime time.Duration) {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(processingTime)
    fmt.Fprintf(w, "Hello from Server %s (Processing Time: %s)", port, processingTime)
  })

  log.Printf("Server started on port %s", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}

func sendRequest(url string, interval time.Duration) {
  ticker := time.NewTicker(interval)
  defer ticker.Stop()

  for range ticker.C {
    resp, err := http.Get(url)
    if err != nil {
      log.Printf("Error sending request: %v", err)
      continue
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Printf("Error reading request: %v", err)
      continue
    }

    log.Printf("Response: %s", string(body))
  }
}

func main() {
  go startServer("8080", 100*time.Millisecond)
  go startServer("8081", 1*time.Second)
  go startServer("8082", 3*time.Second)

  time.Sleep(1 * time.Second)

  go sendRequest("http//localhost:8000", 500*time.Millisecond)

  testDuration := 10 * time.Second)
  time.Sleep(testDuration)
}


