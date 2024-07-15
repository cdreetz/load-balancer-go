package balancer

import (
  "sync/atomic"
)

type RoundRobin struct {
  Backends    []*Backend
  Current     uint64
}

func (r *RoundRobin) NextIndex() int {
  return int(atomic.AddUint64(&r.Current, 1) % uint64(len(r.Backends)))
}

func (r *RoundRobin) GetBackend() *Backend {
  next := r.NextIndex()
  l := len(r.Backends) + next
  for i := next; i < l; i++ {
    idx := i % len(r.Backends)
    if r.Backends[idx].Alive {
      if i != next {
        atomic.StoreUint64(&r.Current, uint64(idx))
      }
      return r.Backends[idx]
    }
  }
  return nil
}
