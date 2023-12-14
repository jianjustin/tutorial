package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/jianjustin/add/handler"
	pb "github.com/jianjustin/add/proto"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service      = "add"
	version      = "latest"
	etcd_address = "localhost:2379"
	address      = ":60003"
)

func main() {
	//etcd registry
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(etcd_address),
	)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(address))),
		micro.Client(grpcc.NewClient()),
		micro.Registry(etcdRegistry),
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
