package main

import (
	"context"
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

	client := NewAddClient(cc)
	_, v, err := client.Add(context.Background(), int64(42))
	if err != nil {
		log.Fatalf("unable to Test: %+v", err)
	}
	log.Println(v)
}
