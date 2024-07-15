package server

import (
  "fmt"
  "log"
  "net/http"

  "github.com/cdreetz/load-balancer-go/balancer"
  "github.com/cdreetz/load-balancer-go/config"
)

func Run(cfg *config.Config, balancer balancer.Balancer) {
  server := http.Server{
    Addr:     fmt.Sprintf(":%s", cfg.Port),
    Handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      backend := balancer.GetBackend()
      if backend != nil {
        backend.ReverseProxy.ServeHTTP(w, r)
        return
      }
      http.Error(w, "No available backends", http.StatusServiceUnavailable)
    }),
  }

  log.Printf("Load balancer started on port %s", cfg.Port)
  log.Fatal(server.ListenAndServe())
}
