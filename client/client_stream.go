package main

import (
	"context"
	"log"
	"time"

	pb "github.com/YatharthJangid/grpc_demo/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("couldnt send names :%v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name:%s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil {
		log.Fatalf("Error whiole receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
