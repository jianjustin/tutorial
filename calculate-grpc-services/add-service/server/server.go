package main

import (
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/register"
	"go.guide/add-grpc-service/service"
	"go.guide/add-grpc-service/transport"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	hostPort string = "localhost:8002"
)

func main() {
	server := grpc.NewServer()
	sc, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatalf("unable to listen: %+v", err)
	}
	defer server.GracefulStop()

	pb2.RegisterAddServiceServer(server, transport.MakeAddGRPCServer(service.NewAddService()))

	r := register.GetEtcdRegister()
	if r == nil {
		fmt.Println("get register client failed")
		return
	}
	err = r.Register(etcdv3.Service{Key: "/services/add/", Value: hostPort})
	if err != nil {
		fmt.Println("register service failed")
		return
	}
	defer r.Deregister(etcdv3.Service{Key: "/services/add/", Value: hostPort})

	_ = server.Serve(sc)
}
