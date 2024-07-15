# Load Balancer in Go


- `main.go` is the applications entry point
- `/balancer` contains all the core logic related to load balancing
  - `backend.go` defines the Backend struct and related methods
  - `health_check.go` implement sthe health check functionality
  - `round_robin.go` implements the round robin load balancing algorithm
  - `balancer.go` defines the Balancer interface and provides a factory function to create a load balancer with the specified algorithm
- `/config` contains the configuration-related code
  - `config.go` defines the Config struct and provides a function to load the configuration from environment variables
- `/server` contains the server implementation
  - `server.go` implements the load balancer server that receives requests and forwards them to the selected backend


  ## Getting Started

  1. Clone the repository:

  `git clone https://github.com/cdreetz/load-balancer-go.git`

  2. Navigate to the project directory:

  `cd load-balancer-go`

  3. Set up the required environment variables:
  
  - PORT: The port on which the load balancer will run (default: 8000)
  - BACKENDS: A comma-separated list of backend server URLs (default: localhost 8080 and 8081)
  - ALGORITHM: The load balancing algorithm to use (default: round-robin)



