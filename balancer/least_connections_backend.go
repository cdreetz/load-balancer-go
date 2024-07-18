package balancer

import "sync/atomic"

type LeastConnectionsBackend struct {
  *Backend
  Connections int64
}

func NewLeastConnectionsBackend(backend *Backend) *LeastConnectionsBackends {
  return &LeastConnectionsBackend{
    Backend:      backend,
    Connections:  0,
  }
}
