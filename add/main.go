package main

import (
	"github.com/jianjustin/add/handler"
	pb "github.com/jianjustin/add/proto"
	"go-micro.dev/v4/server"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "add"
	version = "latest"
	address = ":60003"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(address))),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterAddHandler(srv.Server(), new(handler.Add)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
