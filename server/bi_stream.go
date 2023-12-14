package main

import (
	"io"
	"log"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func (s *helloServer) SayHelloBidirectionalStream(stream pb.GreetService_SayHelloBidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Finished processing stream \n")
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Processing %v \n", req.Name)
		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + req.Name,
		}); err != nil {
			return err
		}
		//time.Sleep(1 * time.Second)
	}
}
