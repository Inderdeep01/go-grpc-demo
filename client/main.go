package main

import (
	"log"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) //WithInsecure()) //insecure.NewCredentials()
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	names := &pb.NameList{
		Names: []string{"Inderdeep", "Singh", "Sidhu", "Golang", "Alice", "Bob"},
	}
	if client == nil {
		log.Fatalf("client is nil")
	} else {
		log.Printf("client %v\n", client)
		//callSayHello(client) //-> call to unary function
		//callSayHelloServerStream(client, names) //-> call to server stream function
		//callSayHelloClientStream(client, names) //-> call to client stream function
		callSayHelloBidirectionalStream(client, names) //-> call to bidirectional stream function
	}
}
