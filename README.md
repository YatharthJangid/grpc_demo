# gRPC Microservice in Go

A complete gRPC service demonstrating all four RPC communication patterns: unary, server streaming, client streaming, and bidirectional streaming.

## Features

- Unary RPC - Simple request-response communication
- Server Streaming - Server sends multiple responses for one client request
- Client Streaming - Client sends multiple requests, server responds once
- Bidirectional Streaming - Real-time two-way communication with concurrent message exchange
- Context-based timeout management
- Protocol Buffers for efficient serialization

## Tech Stack

- Go 1.21+
- gRPC
- Protocol Buffers (proto3)
- Context package

## Project Structure

    grpc_demo/
    ├── server/
    │   ├── main.go
    │   ├── unary.go
    │   ├── server_stream.go
    │   ├── client_stream.go
    │   └── bi_stream.go
    ├── client/
    │   ├── main.go
    │   ├── unary.go
    │   ├── server_stream.go
    │   ├── client_stream.go
    │   └── bi_stream.go
    └── proto/
        ├── greet.proto
        ├── greet.pb.go
        └── greet_grpc.pb.go

## Setup

Clone the repository:

    git clone https://github.com/YatharthJangid/grpc_demo.git
    cd grpc_demo

Install dependencies:

    go mod download

Generate protobuf code if needed:

    protoc --go_out=. --go-grpc_out=. proto/greet.proto

## Running

Start server:

    cd server
    go run *.go

Run client in another terminal:

    cd client
    go run *.go

## RPC Methods

**SayHello** - Unary RPC with simple request-response

**SayHelloServerStreaming** - Server streams multiple responses

**SayHelloClientStreaming** - Client streams multiple requests

**SayHelloBiDirectionalStreaming** - Real-time bidirectional streaming

## What I Learned

- gRPC service implementation with all RPC patterns
- Protocol Buffer schema design
- Go concurrency with goroutines and channels
- Context management for timeouts
- Stream lifecycle management

## Author

Yatharth Jangid

GitHub: github.com/YatharthJangid

LinkedIn: linkedin.com/in/yatharth-jangid
