package main

import (
	"github.com/jianjustin/sub/handler"
	pb "github.com/jianjustin/sub/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "sub"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterSubHandler(srv.Server(), new(handler.Sub)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
