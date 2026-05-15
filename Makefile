.PHONY: all build run-server run-client clean test generate

all: build

build:
	go build -o bin/server ./server/...
	go build -o bin/client ./client/...

run-server:
	go run ./server/*.go

run-client:
	go run ./client/*.go

clean:
	rm -rf bin/

test:
	go test -v ./...

generate:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto
