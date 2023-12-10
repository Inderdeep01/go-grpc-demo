package main

import (
	"log"
	"net"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Started server at %v\n", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
