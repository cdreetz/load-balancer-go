package balancer

import (
  "net/http/httputil"
  "net/url"
)

type Backend struct {
  URL           *url.URL
  Alive         bool
  ReverseProxy  *httputil.ReverseProxy
}

func NewBackend(rawURL string) *Backend {
  url, _ := url.Parse(rawURL)
  return &Backend{
    URL:            url,
    Alive:          true,
    ReverseProxy:   httputil.NewSingleHostReverseProxy(url),
  }
}
