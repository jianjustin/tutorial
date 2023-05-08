package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	helloworldpb "go.guide/grpc/06-nameResolving/proto/helloworld"
	"google.golang.org/grpc"
)

var (
	addrs = []string{":50051", ":50052"}
)

type server struct {
	addr string
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: fmt.Sprintf("%s world in %s\n", in.Name, s.addr)}, nil
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	for _, port := range addrs {
		wg.Add(1)
		go func(port string) {
			defer wg.Done()
			lis, err := net.Listen("tcp", fmt.Sprintf("%s", port))
			if err != nil {
				log.Fatalln("Failed to listen:", err)
			}

			// Create a gRPC server object
			s := grpc.NewServer()
			// Attach the Greeter service to the server
			helloworldpb.RegisterGreeterServer(s, &server{addr: fmt.Sprintf("0.0.0.0%s", port)})
			// Serve gRPC server
			log.Printf("Serving gRPC on 0.0.0.0%s\n", port)
			log.Fatalln(s.Serve(lis))
		}(port)
	}
	wg.Wait()
}
