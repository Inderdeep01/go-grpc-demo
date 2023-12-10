package main

import (
	"context"
	"log"
	"time"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	start := time.Now() // Record the time before the request

	res, err := client.SayHello(ctx, &pb.NoParam{})
	latency := time.Since(start) // Calculate the latency
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %v", res.Message)
	log.Printf("Latency: %v", latency)
}
