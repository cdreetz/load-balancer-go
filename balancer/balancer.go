package balancer

type Balancer interface {
  GetBackend() *Backend
}

func New(algorithm string, backends []*Backend) Balancer {
  switch algorithm {
  case "round-robin":
    return &RoundRobin{Backends: backends}
  default:
    return &RoundRobin{Backends: backends}
  }
}
