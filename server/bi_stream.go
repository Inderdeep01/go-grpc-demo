package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"go.opentelemetry.io/otel/trace"
)

func (s *helloServer) SayHelloBidirectionalStream(stream pb.GreetService_SayHelloBidirectionalStreamServer) error {
	//log.Printf("stream context %v", stream.Context())
	/* _, span := tracer.Start(stream.Context(), "Bidirectional stream request reached server")
	defer span.End() */
	span := trace.SpanFromContext(stream.Context())
	span.AddEvent("Starting Stream")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			span.AddEvent("Finished processing stream on server side")
			log.Printf("Finished processing stream \n")
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Processing %v \n", req.Name)
		span.AddEvent(fmt.Sprintf("Processing %v \n", req.Name))
		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + req.Name,
		}); err != nil {
			return err
		}
		//time.Sleep(1 * time.Second)
	}
}
