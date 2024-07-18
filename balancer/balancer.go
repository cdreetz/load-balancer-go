package balancer

type Balancer interface {
  GetBackend() *Backend
}

func New(algorithm string, backends []*Backend) Balancer {
  switch algorithm {
  case "round-robin":
    return &RoundRobin{Backends: backends}
  case "least-connections":
    leastConnectionsBackends := make([]*LeastConnectionsBackend, len(backends))
    for i, backend := range backends {
      leastConnectionsBackends[i] = NewLeastConnectionsBackend(backend)
    }
    return &LeastConnections{Backends: leastConnectionsBackends}
  default:
    return &RoundRobin{Backends: backends}
  }
}
