package main

import (
  "github.com/cdreetz/load-balancer-go/balancer"
  "github.com/cdreetz/load-balancer-go/config"
  "github.com/cdreetz/load-balancer-go/server"
)

func main() {
  cfg := config.Load()
  backends := make([]*balancer.Backend, len(cfg.Backends))
  for i, url := range cfg.Backends {
    backends[i] = balancer.NewBackend(url)
    go backends[i].HealthCheck()
  }
  lb := balancer.New(cfg.Algorithm, backends)
  server.Run(cfg, lb)
}
