package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	a "go.guide/microservices-app-01/proto/a"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	restful = flag.Int("restful", 8080, "the port to restful serve on")
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

	// Serve gRPC server
	log.Printf("Serving gRPC on 0.0.0.0:%d\n", *port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", *port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = a.RegisterAHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *restful),
		Handler: gwmux,
	}

	log.Println(fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0::%d", *restful))
	log.Fatalln(gwServer.ListenAndServe())
}
