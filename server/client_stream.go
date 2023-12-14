package main

import (
	"io"
	"log"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.MessageList{Messages: messages})
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received request from client: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
