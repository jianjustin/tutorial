package main

import (
	"storage/handler"
	"storage/middleware"
	pb "storage/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "storage"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Address(":50000"),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	db := middleware.NewPostgresInstance()

	// Register handler
	if err := pb.RegisterStorageHandler(srv.Server(), handler.NewStorageHandler(db)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
