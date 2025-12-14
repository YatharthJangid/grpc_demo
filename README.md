
```markdown
# gRPC Microservice in Go

A complete gRPC service demonstrating all four RPC communication patterns: unary, server streaming, client streaming, and bidirectional streaming.

## 🚀 Features

- **Unary RPC**: Simple request-response communication
- **Server Streaming**: Server sends multiple responses for one client request
- **Client Streaming**: Client sends multiple requests, server responds once
- **Bidirectional Streaming**: Real-time two-way communication with concurrent message exchange
- Context-based timeout management
- Protocol Buffers for efficient serialization
- Goroutines and channels for concurrent streaming operations

## 🛠️ Tech Stack

- **Go 1.21+**
- **gRPC** - High-performance RPC framework
- **Protocol Buffers (proto3)** - Interface definition and serialization
- **Context** - Request lifecycle management

## 📁 Project Structure

```
grpc_demo/
├── server/
│   ├── main.go              # Server initialization and routing
│   ├── unary.go             # Unary RPC handler
│   ├── server_stream.go     # Server streaming handler
│   ├── client_stream.go     # Client streaming handler
│   └── bi_stream.go         # Bidirectional streaming handler
├── client/
│   ├── main.go              # Client connection setup
│   ├── unary.go             # Unary RPC client call
│   ├── server_stream.go     # Server streaming client
│   ├── client_stream.go     # Client streaming implementation
│   └── bi_stream.go         # Bidirectional streaming client
└── proto/
    ├── greet.proto          # Protocol Buffer service definition
    ├── greet.pb.go          # Generated message code
    └── greet_grpc.pb.go     # Generated gRPC service code
```

## 📋 Prerequisites

- Go 1.21 or higher
- Protocol Buffer compiler (`protoc`)
- gRPC Go plugins

### Install Dependencies

```
# Install protoc compiler
# Windows: Download from https://github.com/protocolbuffers/protobuf/releases
# Linux: sudo apt install protobuf-compiler
# Mac: brew install protobuf

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 🔧 Setup & Installation

1. **Clone the repository**
```
git clone https://github.com/YatharthJangid/grpc_demo.git
cd grpc_demo
```

2. **Install Go dependencies**
```
go mod download
```

3. **Generate protobuf code** (only if you modify `greet.proto`)
```
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

## ▶️ Running the Application

### Start the Server

Open a terminal and run:

```
cd server
go run *.go
```

**Expected output:**
```
server started at [::]:8080
```

### Run the Client

Open another terminal and run:

```
cd client
go run *.go
```

You can switch between different RPC patterns by uncommenting the desired method in `client/main.go`:

```
// Uncomment one of these:
callSayHello(client)                          // Unary
callSayHelloServerStreaming(client, names)    // Server streaming
callSayHelloClientStreaming(client, names)    // Client streaming
callHelloBiDirectionalStream(client, names)   // Bidirectional (default)
```

## 📡 RPC Methods Explained

### 1. **SayHello** (Unary RPC)

**Pattern:** Client → Server → Client

Simple request-response. Client sends one request, server sends one response.

**Example:**
```
Client: "Hello!"
Server: "Hello back!"
```

### 2. **SayHelloServerStreaming** (Server Streaming)

**Pattern:** Client → Server → → → Client

Client sends a list of names, server streams back greetings one by one.

**Example:**
```
Client: ["Yatharth", "Alice", "Bob"]
Server: "Hello Yatharth" → "Hello Alice" → "Hello Bob"
```

### 3. **SayHelloClientStreaming** (Client Streaming)

**Pattern:** Client → → → Server → Client

Client streams multiple names, server collects all and responds once.

**Example:**
```
Client: "Yatharth" → "Alice" → "Bob" (closes)
Server: "Hello Yatharth, Alice, Bob"
```

### 4. **SayHelloBiDirectionalStreaming** (Bidirectional Streaming)

**Pattern:** Client ↔ ↔ ↔ Server

Both client and server send messages independently and concurrently. Like a chat application.

**Implementation highlights:**
- Uses goroutines for concurrent send/receive operations
- Channel-based synchronization with `waitc`
- `stream.CloseSend()` signals client finished sending
- `io.EOF` signals server finished responding

**Example:**
```
Client sends: "Yatharth" → Server responds: "Hello Yatharth"
Client sends: "Alice"    → Server responds: "Hello Alice"
Client sends: "Bob"      → Server responds: "Hello Bob"
(All happen concurrently)
```

## 📊 Sample Output

**Bidirectional Streaming Output:**

```
Client:
bidirn streaming has started
message:"Hello Yatharth"
message:"Hello Alice"
message:"Hello Bob"
Bidirn streaming finished

Server:
server started at [::]:8080
Got req with name: Yatharth
Got req with name: Alice
Got req with name: Bob
```

## 🎓 Key Concepts Learned

### gRPC Fundamentals
- Protocol Buffer schema design
- Service definition and code generation
- Four RPC communication patterns
- Stream lifecycle management

### Go Concurrency
- Goroutines for parallel operations
- Channels for synchronization (`chan struct{}`)
- Context management (timeouts, cancellation)
- `defer` for resource cleanup

### Distributed Systems
- Client-server architecture
- Stream-based communication
- Error handling in network calls
- Connection management and cleanup

## 🔮 Future Enhancements

- [ ] Add JWT-based authentication
- [ ] Implement gRPC interceptors for logging/metrics
- [ ] Add unit tests and integration tests
- [ ] TLS/SSL encryption for secure communication
- [ ] Load balancing across multiple servers
- [ ] Deploy to cloud (AWS Lambda / Google Cloud Run)
- [ ] Add health check endpoints
- [ ] Implement rate limiting
- [ ] Add OpenTelemetry tracing

## 🐛 Troubleshooting

**Issue:** `protoc-gen-go: program not found`
```
# Add Go bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

**Issue:** `connection refused`
- Make sure server is running first
- Check port 8080 is not in use
- Verify server logs show "server started"

**Issue:** `Error while streaming`
- Ensure method names match between client and server
- Check server implementation has all RPC methods
- Verify proto file is properly compiled

## 📝 Protocol Buffer Definition

```
service greet_service {
    rpc SayHello (NoParam) returns (HelloResponse);
    rpc SayHelloServerStreaming (NamesList) returns (stream HelloResponse);
    rpc SayHelloClientStreaming (stream HelloRequest) returns (MessagesList);
    rpc SayHelloBiDirectionalStreaming (stream HelloRequest) returns (stream HelloResponse);
}
```

## 👤 Author

**Yatharth Jangid**  
- GitHub: [@YatharthJangid](https://github.com/YatharthJangid)
- LinkedIn: [yatharth-jangid](https://linkedin.com/in/yatharth-jangid)
- Email: jangidyatharth5@gmail.com

## 📄 License

This project is open source and available for educational purposes.

---

⭐ **Star this repo if you found it helpful!**

Built while learning distributed systems and microservices architecture.
```