package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcd_pkg "go.guide/grpc/08-etcd/pkg"
	"go.guide/grpc/08-etcd/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

var (
	addrs = []string{":50051", ":50052"}
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	Addr string
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName() + s.Addr}, nil
}

func main() {

	client, err := clientv3.New(clientv3.Config{Endpoints: []string{"0.0.0.0:2379"}, DialTimeout: time.Second * 5})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()

			go func(addr string) {
				err := etcd_pkg.Register(ctx, client, "hello-world-service", "localhost"+addr)
				if err != nil {
					log.Fatal(err)
				}
			}(addr)

			lis, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			helloworld.RegisterGreeterServer(s, &server{Addr: fmt.Sprintf("0.0.0.0%s", addr)})
			log.Printf("server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}(addr)

	}
	wg.Wait()
}
