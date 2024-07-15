package balancer

import (
  "log"
  "net/http"
  "time"
)

func (b *Backend) HealthCheck() {
  client := &http.Client{
    Timeout: time.Second * 10,
  }

  ticker := time.NewTicker(time.Second * 10)
  for range ticker.C {
    resp, err := client.Get(b.URL.String())
    if err != nil || resp.StatusCode != http.StatusOK {
      b.Alive = false
      log.Printf("Backend %s is down", b.URL.String())
      continue
    }
    b.Alive = true
    log.Printf("Backend %s is up", b.URL.String())
  }
}
