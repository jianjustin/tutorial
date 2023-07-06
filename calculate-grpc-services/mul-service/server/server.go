package main

import (
	"go.guide/mul-grpc-service/pb"
	"go.guide/mul-grpc-service/service"
	"go.guide/mul-grpc-service/transport"
	"google.golang.org/grpc"
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
	_ = server.Serve(sc)
}
