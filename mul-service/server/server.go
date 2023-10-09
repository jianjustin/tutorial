package main

import (
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
	"jianjustin/mul-grpc-service/pb"
	"jianjustin/mul-grpc-service/register"
	"jianjustin/mul-grpc-service/service"
	"jianjustin/mul-grpc-service/transport"
	"log"
	"net"
)

const (
	hostPort string = "localhost:8003"
)

func main() {
	server := grpc.NewServer()
	sc, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatalf("unable to listen: %+v", err)
	}
	defer server.GracefulStop()

	pb.RegisterMulServiceServer(server, transport.MakeMulGRPCServer(service.NewMulService()))

	r := register.GetEtcdRegister()
	if r == nil {
		fmt.Println("get register client failed")
		return
	}
	err = r.Register(etcdv3.Service{Key: "/services/mul/", Value: hostPort})
	if err != nil {
		fmt.Println("register service failed")
		return
	}
	defer r.Deregister(etcdv3.Service{Key: "/services/mul/", Value: hostPort})

	_ = server.Serve(sc)
}
