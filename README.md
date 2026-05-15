# gRPC Microservice in Go 🚀

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-244c5a?style=for-the-badge&logo=grpc&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

A complete and production-ready gRPC service demonstrating all four RPC communication patterns in Go: unary, server streaming, client streaming, and bidirectional streaming.

## ✨ Features

- **Unary RPC** - Simple request-response communication.
- **Server Streaming** - Server sends multiple responses for one client request.
- **Client Streaming** - Client sends multiple requests, server responds once.
- **Bidirectional Streaming** - Real-time two-way communication with concurrent message exchange.
- **Interceptors** - Basic interceptor implemented for logging incoming requests.
- **Containerization** - Ready to be deployed using Docker.
- **CI/CD** - Automated builds and testing with GitHub Actions.
- **Context-based timeout management**
- **Protocol Buffers for efficient serialization**

## 🛠️ Tech Stack

- **Language:** Go 1.21+
- **RPC Framework:** gRPC
- **Serialization:** Protocol Buffers (proto3)

## 📂 Project Structure

```text
grpc_demo/
├── client/                 # gRPC Client implementation
│   ├── bi_stream.go        # Bidirectional streaming client
│   ├── client_stream.go    # Client streaming client
│   ├── main.go             # Client entry point
│   ├── server_stream.go    # Server streaming client
│   └── unary.go            # Unary client
├── proto/                  # Protocol Buffers definitions
│   ├── greet.proto         # Service schema
│   ├── greet.pb.go         # Generated protobuf structs
│   └── greet_grpc.pb.go    # Generated gRPC code
├── server/                 # gRPC Server implementation
│   ├── bi_stream.go        # Bidirectional streaming handler
│   ├── client_stream.go    # Client streaming handler
│   ├── main.go             # Server entry point with interceptors
│   ├── server_stream.go    # Server streaming handler
│   ├── server_test.go      # Unit tests
│   └── unary.go            # Unary handler
├── Dockerfile              # Containerization for the server
├── Makefile                # Build and run commands
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Protocol Buffer Compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

### Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/YatharthJangid/grpc_demo.git
   cd grpc_demo
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **(Optional) Re-generate protobuf code:**
   ```bash
   make generate
   ```

### 🏃 Running the Application

Using the provided `Makefile`:

**Start the Server:**
```bash
make run-server
```

**Run the Client (in another terminal):**
```bash
make run-client
```

### 🐳 Running with Docker

You can easily build and run the gRPC server using Docker:

```bash
docker build -t grpc-demo-server .
docker run -p 8080:8080 grpc-demo-server
```

## 🧪 Testing

Run the included unit tests using:

```bash
make test
```

## 📚 RPC Methods Implemented

| Method | Type | Description |
|--------|------|-------------|
| `SayHello` | Unary | Simple request-response. Client sends an empty param and gets a "Hello" response. |
| `SayHelloServerStreaming` | Server Streaming | Client sends a list of names, server streams a greeting for each name. |
| `SayHelloClientStreaming` | Client Streaming | Client streams multiple names, server replies with a single summary message. |
| `SayHelloBiDirectionalStreaming` | Bidirectional | Client and Server exchange greetings dynamically in real-time. |

## 🎓 What I Learned

- Designing and building gRPC services with all four RPC communication patterns.
- Protocol Buffer schema design and Go code generation.
- Go concurrency paradigms using `goroutines` and `channels`.
- Request lifecycle and Context management (handling timeouts/cancellations).
- Adding gRPC Interceptors for middleware-like functionality (logging, auth).
- Containerizing Go applications efficiently using multi-stage Docker builds.

## ✍️ Author

**Yatharth Jangid**
- GitHub: [YatharthJangid](https://github.com/YatharthJangid)
- LinkedIn: [yatharth-jangid](https://linkedin.com/in/yatharth-jangid)
