package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  port := os.Args[1]
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from backend server on port %s!", port)
  })
  log.Printf("Backed server started on %s", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
