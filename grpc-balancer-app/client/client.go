package main

import (
	"context"
	"fmt"
	"go.guide/grpc-balancer-app/balancer"
	pb "go.guide/grpc-balancer-app/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

func main() {
	r := balancer.NewResolver("localhost:2378")
	resolver.Register(r)

	const grpcServiceConfig = `{"loadBalancingPolicy":"round_robin"}`
	conn, err := grpc.Dial(r.Scheme()+"://author/project/test", grpc.WithDefaultServiceConfig(grpcServiceConfig), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewGreeterClient(conn)

	for {
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "abc"}, grpc.FailFast(true))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}

		<-time.After(time.Second)
	}
}
