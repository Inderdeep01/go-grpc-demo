package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloBidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("failed to call SayHelloBidirectionalStream: %v", err)
	}
	log.Print("Starting stream \n")
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send a request to server: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive a response from server: %v", err)
		}
		log.Printf("Response from server: %v\n", res.Message)
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
}
