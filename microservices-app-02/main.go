package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	b "go.guide/microservices-app-02/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	b.UnimplementedBServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddBPreffix(ctx context.Context, in *b.BRequest) (*b.BReply, error) {
	//log.Printf("Received: %v", in.GetName())

	return &b.BReply{Res: fmt.Sprintf("%s02", in.GetName())}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	b.RegisterBServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
