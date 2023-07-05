package main

import (
	"context"
	"github.com/go-kit/kit/sd/etcdv3"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.guide/mul-service/endpoint"
	"go.guide/mul-service/service"
	"log"
	"net/http"
	"time"
)

func main() {
	svc := service.BaseMulService{}

	mulHandler := httptransport.NewServer(
		endpoint.MakeMulEndpoint(svc),
		endpoint.DecodeMulRequest,
		endpoint.EncodeResponse,
	)

	mulAfterAddHandler := httptransport.NewServer(
		endpoint.MakeMulAfterAddEndpoint(svc),
		endpoint.DecodeMulRequest,
		endpoint.EncodeResponse,
	)

	http.Handle("/mul", mulHandler)
	http.Handle("/mulAfterAdd", mulAfterAddHandler)

	log.Println("Server started on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	registerService()
}

func registerService() {
	prefix := "/services/add/"
	instance := ":8081"
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
