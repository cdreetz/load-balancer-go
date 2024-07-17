package balancer

import (
  "sync"
)

type LeastConnections struct {
  Backends    []*Backend
  connections []int
  mu          sync.Mutex
}

func (l *LeastConnections) initConnections() {
  l.connections = make([]int, len(l.Backends))
}

func (l *LeastConnections) GetBackend() *Backend {
  l.mu.Lock()
  defer l.mu.Unlock()
  var least int
  var idx int
  for i, c := range l.connections {
    if c < least || least == 0 {
      least = c
      idx = i
    }
  }
  if !l.Backends[idx].Alive {
    return nil
  }
  l.connections[idx]++
  return l.Backends[idx]
}

func (l *LeastConnections) Done(backend *Backend) {
  l.mu.Lock()
  defer l.mu.Unlock()
  for i, b := range l.Backends {
    if b == backend {
      l.connections[i]--
      return
    }
  }
}
