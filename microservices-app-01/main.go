package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	a "go.guide/microservices-app-01/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	a.UnimplementedAServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddAPreffix(ctx context.Context, in *a.ARequest) (*a.AReply, error) {
	//log.Printf("Received: %v", in.GetName())
	
	return &a.AReply{Res: fmt.Sprintf("%s01", in.GetName())}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	a.RegisterAServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
