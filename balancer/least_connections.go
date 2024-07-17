package balancer

import (
  "sync"
)

type LeastConnections struct {
  Backends    []*Backend
  mutex          sync.Mutex
}

func (l *LeastConnections) GetBackend() *Backend {
  l.mutex.Lock()
  defer l.mutex.Unlock()

  var leastConnections int
  var leastConnectionsBackend *Backend

  for _, backend := range l.Backends {
    if backend.Alive {
      if leastConnectionsBackend == nil || backend.Connections < leastConnections {
        leastConnections = backend.Connections
        leastConnectionsBackend = backend
      }
    }
  }

  if leastConnectionsBackend != nil {
    leastConnectionsBackend.Connections++
  }

  return leastConnectionsBackend
}
