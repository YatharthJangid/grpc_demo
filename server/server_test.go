package main

import (
	"context"
	"testing"

	pb "github.com/YatharthJangid/grpc_demo/proto"
)

func TestSayHello(t *testing.T) {
	server := &helloServer{}

	req := &pb.NoParam{}
	res, err := server.SayHello(context.Background(), req)

	if err != nil {
		t.Errorf("SayHello failed with error: %v", err)
	}

	expected := "Hello"
	if res.Message != expected {
		t.Errorf("Expected %v, got %v", expected, res.Message)
	}
}
