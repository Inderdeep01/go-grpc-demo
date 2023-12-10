package main

import (
	"log"
	"time"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("Received request from client: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
