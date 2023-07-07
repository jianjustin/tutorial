package main

import (
	"fmt"
	"github.com/go-kit/log"
	log2 "go.guide/add-grpc-service/log"
	"go.guide/add-grpc-service/middleware"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/register"
	"go.guide/add-grpc-service/service"
	"go.guide/add-grpc-service/transport"
	"google.golang.org/grpc"
	"net"
	"os"
)

const ()

func main() {
	server := grpc.NewServer()
	logger := log.With(log2.InitLogger(os.Stdout))
	sc, err := net.Listen("tcp", middleware.HostPort)
	if err != nil {
		log.With(logger, "level", "error").Log("msg", fmt.Sprintf("unable to listen %s", err))
	}
	defer server.GracefulStop()

	svc := service.NewAddService()
	svc = middleware.LoggingAddServiceMiddleware(logger)(svc)
	svc = middleware.EtcdRegisterAddServiceMiddleware(register.GetEtcdRegister(), logger)(svc)

	pb2.RegisterAddServiceServer(server, transport.MakeAddGRPCServer(service.NewAddService()))

	log.With(logger, "level", "info").Log("msg", fmt.Sprintf("grpc server start at %s", middleware.HostPort))
	_ = server.Serve(sc)
}
