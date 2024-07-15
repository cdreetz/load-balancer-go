package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "sync/atomic"
)

type Backend struct {
  URL           *url.URL
  Alive         bool
  ReverseProxy  *httputil.ReverseProxy
}

type ServerPool struct {
  Backends       []*Backend
  Current        uint64
}

func (s *ServerPool) NextIndex() int {
  return int(atomic.AddUint64(&s.Current, uint64(1)) % uint64(len(s.Backends)))
}

func (s *ServerPool) GetNextBackend() *Backend {
  next := s.NextIndex()
  l := len(s.Backends) + next
  for i := next; i < l; i++ {
    idx := i % len(s.Backends)
    if s.Backends[idx].Alive {
      if i != next {
        atomic.StoreUint64(&s.Current, uint64(idx))
      }
      return s.Backends[idx]
    }
  }
  return nil
}

func (s *ServerPool) HealthCheck() {
  for _, b := range s.Backends {
    status := "up"
    resp, err := http.Head(b.URL.String())
    if err != nil || resp.StatusCode != 200 {
      status = "down"
      b.Alive = false
    } else {
      b.Alive = true
    }
    log.Printf("%s [%s]\n", b.URL, status)
  }
}

func (s *ServerPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  backend := s.GetNextBackend()
  if backend != nil {
    backend.ReverseProxy.ServeHTTP(w, r)
    return
  }
  http.Error(w, "No available backends", http.StatusServiceUnavailable)
}

func main() {
  backends := []*Backend{
    {URL: &url.URL{Scheme: "http", Host: "localhost:8080"}},
    {URL: &url.URL{Scheme: "http", Host: "localhost:8081"}},
    {URL: &url.URL{Scheme: "http", Host: "localhost:8082"}},
  }
  serverPool := &ServerPool{}
  for _, b := range backends {
    serverPool.Backends = append(serverPool.Backends, &Backend{
      URL:           b.URL,
      Alive:         true,
      ReverseProxy:  httputil.NewSingleHostReverseProxy(b.URL),
    })
    serverPool.HealthCheck()
  }
  server := http.Server{
    Addr:      "localhost:8000",
    Handler:   serverPool,
  }
  fmt.Printf("Load Balancer started at %s\n", server.Addr)
  log.Fatal(server.ListenAndServe())
}


