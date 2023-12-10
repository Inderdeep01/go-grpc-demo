package main

import (
	"context"
	"io"
	"log"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Starting stream \n")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("failed to call SayHelloServerStream: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive a message from server: %v", err)
		}
		log.Printf("Response from server: %v", msg)
	}
	log.Printf("Finished stream \n")
}
