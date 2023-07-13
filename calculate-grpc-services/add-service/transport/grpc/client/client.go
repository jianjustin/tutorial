package main

import (
	"context"
	transportgrpc "go.guide/add-grpc-service/transport/grpc"
	"google.golang.org/grpc"
	"log"
)

const (
	hostPort string = "localhost:8002"
)

func main() {
	cc, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to Dial: %+v", err)
	}

	client := transportgrpc.NewAddClient(cc)
	_, v, err := client.Add(context.Background(), int64(42))
	if err != nil {
		log.Fatalf("unable to Test: %+v", err)
	}
	log.Println(v)

	client1 := transportgrpc.NewAddAfterMulClient(cc)
	_, v, err = client1.AddAfterMul(context.Background(), int64(42))
	if err != nil {
		log.Fatalf("unable to Test: %+v", err)
	}
	log.Println(v)
}
