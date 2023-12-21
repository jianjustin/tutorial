package main

import (
	"fmt"
	mgrpc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/etcd"
	mhttp "github.com/go-micro/plugins/v4/server/http"
	"github.com/gorilla/mux"
	"github.com/jianjustin/frontend/config"
	"github.com/jianjustin/frontend/proto/add"
	pb "github.com/jianjustin/frontend/proto/mul"
	"github.com/jianjustin/frontend/proto/sub"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"net/http"

	"go-micro.dev/v4"
)

var (
	service      = "frontend"
	version      = "latest"
	etcd_address = "etcd-service:2379"
)

type frontendServer struct {
	addService add.AddService
	mulService pb.MulService
	subService sub.SubService
}

func main() {
	//etcd registry
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(etcd_address),
	)

	// Create service
	srv := micro.NewService(
		micro.Client(mgrpc.NewClient()),
		micro.Server(mhttp.NewServer()),
		micro.Registry(etcdRegistry),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(config.CurrentConfig.Address),
	)

	client := srv.Client()
	svc := &frontendServer{
		addService: add.NewAddService(config.CurrentConfig.AddService, client),
	}

	r := mux.NewRouter()
	r.HandleFunc("/add", svc.AddHandler).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/mul", svc.MulHandler).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/sub", svc.SubHandler).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

	var handler http.Handler = r
	if err := micro.RegisterHandler(srv.Server(), handler); err != nil {
		logger.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
