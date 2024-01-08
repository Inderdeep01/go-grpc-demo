package main

import (
	"context"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"go.opentelemetry.io/otel/trace"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	/* ctx, span := tracer.Start(ctx, "Unary Call on Server")
	defer span.End() */
	span := trace.SpanFromContext(ctx)
	span.AddEvent("Unary Call reached Server")
	return &pb.HelloResponse{
		Message: "Hello from server",
	}, nil
}
