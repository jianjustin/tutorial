package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/jianjustin/sub/handler"
	"github.com/jianjustin/sub/proto/mul"
	pb "github.com/jianjustin/sub/proto/sub"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service      = "sub"
	version      = "latest"
	address      = ":60002"
	etcd_address = "etcd-service:2379"
)

func main() {
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

	// Initialise  mul service
	handler.MulClient = mul.NewMulService("mul", srv.Client())

	// Register handler
	if err := pb.RegisterSubHandler(srv.Server(), new(handler.Sub)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
