# Load Balancer

## Overview

This project is a high-performance load balancer built in Go.optimized for handling HTTP requests and routing them to backend servers. It leverages worker pools for concurrent processing and uses a custom queue to manage incoming traffic efficiently. The load balancer supports both Layer 7 (Application Layer) and Layer 4 (Transport Layer) load balancing, enabling seamless handling of HTTP/HTTPS traffic. It also includes caching mechanisms, such as in-memory caching and Redis caching, to ensure reliability and improved performance.

## Features

- **Concurrent Request Handling**: Implements a worker pool for parallel processing of multiple requests.
- **Request Queueing**: Utilizes Go's thread-safe channels as a queue to organize and dispatch requests to available workers.
- **Worker Management**: Dynamically manages worker availability and handles scenarios where workers are busy.
- **Layer 7 (Application Layer) Load Balancing**: Routes HTTP/HTTPS requests based on URL paths, headers, and other application-level parameters.
- **Layer 4 (Transport Layer) Load Balancing**: Balances traffic based on transport-layer information for lower-latency routing.
- **Caching**: Provides in-memory and Redis caching options to improve performance for repeated requests. Caching is enabled only for GET requests to maintain idempotency.
- **Round-Robin and Weighted Round-Robin Scheduling**: Balances load based on server capacities.
- **Health Checks**: Work in progress (WIP) feature to monitor server health.

## Getting Started

### Prerequisites

- Go (version 1.15 or later recommended)
- Redis (optional, for Redis caching functionality)

### Usage Instructions

1. **Clone the repository**

   ```sh
   git clone https://github.com/swan28pande/Load-Balancer.git
   cd Load-Balancer/load-balancer
   ```

2. **Edit the configuration file**
    Edit the `config.json` according to your needs.
   ```sh
   {
    "port":"8080",
    "level":"L7",
    "strategy":"round-robin",
    "caching":"redis",
    "proto":"tcp",
    "serializer":"none",
    "servers":{
        "http://localhost:8081":3,
        "http://localhost:8082":2
    },
    "maxWorkers":5,
    "cache-ignore":[]
   }
   ```
   
   - `port`: Port on which the loadbalancer service runs
   - `level`: L4/L7 (OSI model layer)
   - `strategy`: Routing strategy, currently supports `round-robin` and `weighted-round-robin`
   - `caching`: Caching mechanism, supports `none`, `baseline` for basic in-memory caching and `redis` for redis caching
   - `servers`: Server URLs and their load handling capacities (number of requests they can handle at once safely)
   - `maxWorkers`: Number of worker threads that will pick requests from channel, it is recommended that the number worker threads be greater than or equal to the total load handling capacity of all backend servers
   - `cache-ignore`: The http methods for which requests will not be cached
   - `proto`: WIP
   - `serializer`: WIP

3. **Dummy Client and Servers**

   - Use as many servers as you need for testing. These dummy servers are just echo servers spitting out random text, the command for starting dummy servers is given below.
   ```sh
   go run main.go 8xxx
   ```
   - Similarly, the test client can be started in a similar manner. There is are two test clients, `concurrent_client` and `client` but it is recommended to use `concurrent_client` instead, due to its ability to emulate multiple clients which is ideal for testing a load-balancer
   ```sh
   go run main.go
   ```

