package balancer

import (
  "sync"
)

type LeastConnections struct {
  Backends       []*LeastConnectionsBackend
  mutex          sync.Mutex
}

func (l *LeastConnections) GetBackend() *Backend {
  l.mutex.Lock()
  defer l.mutex.Unlock()

  var leastConnections int64
  var leastConnectionsBackend *LeastConnectionsBackend

  for _, backend := range l.Backends {
    if backend.Alive {
      if leastConnectionsBackend == nil || backend.Connections < leastConnections {
        leastConnections = backend.Connections
        leastConnectionsBackend = backend
      }
    }
  }

  if leastConnectionsBackend != nil {
    atomic.AddInt64(&leastConnectionsBackend.Connections, 1)
    return leastConnectionsBackend.Backend
  }

  return nil
}
