package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/YatharthJangid/grpc_demo/proto"
)

func callHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("bidirn streaming has started")
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("couldnt send names:%v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			messages, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming")

			}
			log.Println(messages)
		}
		close(waitc)
	}()
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirn streaming finished")

}
