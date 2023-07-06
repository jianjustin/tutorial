package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:8003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to Dial: %+v", err)
	}

	client := NewMulClient(cc)
	_, v, err := client.Mul(context.Background(), int64(42))
	if err != nil {
		log.Fatalf("unable to Test: %+v", err)
	}
	log.Println(v)

	client1 := NewMulAfterAddClient(cc)
	_, v, err = client1.MulAfterAdd(context.Background(), int64(42))
	if err != nil {
		log.Fatalf("unable to Test: %+v", err)
	}
	log.Println(v)
}
