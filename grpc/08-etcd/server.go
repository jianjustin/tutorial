package main

import (
	"context"
	"flag"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcd_pkg "go.guide/grpc/08-etcd/pkg"
	"go.guide/grpc/08-etcd/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var ()

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	port := flag.String("port", ":50051", "The server port")
	service := flag.String("service", "hello-world-service", "service name")

	client, err := clientv3.New(clientv3.Config{Endpoints: []string{"0.0.0.0:2379"}, DialTimeout: time.Second * 5})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := etcd_pkg.Register(ctx, client, *service, "localhost"+*port)
		if err != nil {
			log.Fatal(err)
		}
	}()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
