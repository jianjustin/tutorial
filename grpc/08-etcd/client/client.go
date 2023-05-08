package main

import (
	"context"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcd_pkg "go.guide/grpc/08-etcd/pkg"
	"go.guide/grpc/08-etcd/proto"
	pb "go.guide/grpc/08-etcd/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{"0.0.0.0:2379"}, DialTimeout: time.Second * 5})
	if err != nil {
		log.Fatal(err)
	}

	b := etcd_pkg.NewBuilder(client)

	resolver.Register(b)

	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", etcd_pkg.Scheme, "hello-world-service"), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	i := 1

	for {
		r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: *name})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r.Message)
		}
		i++
		time.Sleep(time.Second)
	}
}
