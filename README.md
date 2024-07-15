# Load Balancer in Go


- `main.go` is the applications entry point
- `/balancer` contains all the core logic related to load balancing
  - `balancer.go` core functionality of and structures of the load balancer
  - `/algorithms` different load balancing strategies
    - `roundrobin.go` implements the round-robin load balancing algorithm
    - `leasconnection.go` implements the least connection algorithm
    - `iphash.go` implements the IP hash-based load balancing



