package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	helloworldpb "go.guide/grpc/06-nameResolving/proto/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	flag.Parse()
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server
	log.Printf("Serving gRPC on 0.0.0.0:%d\n", *port)
	log.Fatalln(s.Serve(lis))
}
