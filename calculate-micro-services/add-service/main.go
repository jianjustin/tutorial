package main

import (
	"context"
	"github.com/go-kit/kit/sd/etcdv3"
	"go.guide/add-service/endpoint"
	"go.guide/add-service/service"
	"log"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := service.BaseAddService{}

	addHandler := httptransport.NewServer(
		endpoint.MakeAddEndpoint(svc),
		endpoint.DecodeAddRequest,
		endpoint.EncodeResponse,
	)

	addAfterMulHandler := httptransport.NewServer(
		endpoint.MakeAddAfterMulEndpoint(svc),
		endpoint.DecodeAddRequest,
		endpoint.EncodeResponse,
	)

	http.Handle("/add", addHandler)
	http.Handle("/addAfterMul", addAfterMulHandler)

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	registerService()
}

func registerService() {
	prefix := "/services/add/"
	instance := ":8080"
	client, err := etcdv3.NewClient(
		context.Background(),
		[]string{"http://127.0.0.1:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	if err != nil {
		log.Printf("unexpected error creating client: %v", err)
	}
	if client == nil {
		log.Printf("expected new Client, got nil")
	}
	err = client.Register(etcdv3.Service{Key: prefix + instance, Value: instance})
	if err != nil {
		log.Printf("unexpected error registering service: %v", err)
	}
	defer client.Deregister(etcdv3.Service{Key: prefix + instance, Value: instance})
}
