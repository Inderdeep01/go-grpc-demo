package main

import (
	"context"
	"log"
	"time"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"go.opentelemetry.io/otel/trace"
)

func callSayHello(ctx context.Context, client pb.GreetServiceClient) {
	/* ctx, span := tracer.Start(ctx, "Sending Request")
	defer span.End() */
	/* carrier := propagation.MapCarrier{}
	propagator := otel.GetTextMapPropagator()
	propagator.Inject(spanContext, carrier) */
	span := trace.SpanFromContext(ctx)
	span.AddEvent("Calling ")

	start := time.Now() // Record the time before the request

	res, err := client.SayHello(ctx, &pb.NoParam{})
	latency := time.Since(start) // Calculate the latency
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %v", res.Message)
	log.Printf("Latency: %v", latency)
	span.AddEvent("Unary Function Call Completed")
}
