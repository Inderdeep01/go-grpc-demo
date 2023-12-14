package main

import (
	"context"
	"log"
	"time"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Starting stream \n")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("failed to call SayHelloClientStream: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send a request to server: %v", err)
		}
		log.Printf("Sent a request to server: %v", name)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Print("Finished stream \n")
	if err != nil {
		log.Fatalf("failed to receive a response from server: %v", err)
	}
	log.Printf("Response from server: %v", res.Messages)
}
