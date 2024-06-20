package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/div-grpc-service/pb"
	"go.guide/div-grpc-service/service"
	"go.guide/div-grpc-service/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"os"
)

var (
	port    = flag.Int("port", 50054, "The server port")
	restful = flag.Int("restful", 8084, "the port to restful serve on")
)

func main() {
	flag.Parse()
	logger := log.NewJSONLogger(os.Stdout)
	g, ctx := errgroup.WithContext(context.Background())

	svc := service.NewDivService()
	endpoints := transport.Endpoints(svc)

	level.Info(logger).Log("msg", fmt.Sprintf("grpc server start at 127.0.0.1:%v", *port))

	g.Go(func() error {
		return ServeGRPC(ctx, &endpoints, fmt.Sprintf("127.0.0.1:%d", *port), log.With(logger, "transport", "GRPC"))
	})

	if err := g.Wait(); err != nil {
		level.Error(logger).Log("error", err)
	}
}

// ServeGRPC serves gRPC requests.
func ServeGRPC(ctx context.Context, endpoints *transport.EndpointsSet, addr string, logger log.Logger) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server := transport.NewGRPCServer(endpoints,
		logger,
		opentracinggo.NoopTracer{},
	)
	grpcServer := grpc.NewServer()
	pb.RegisterDivServiceServer(grpcServer, server)
	level.Info(logger).Log("listen on", addr)

	ch := make(chan error)
	go func() {
		ch <- grpcServer.Serve(listener)
	}()

	select {
	case err := <-ch:
		return fmt.Errorf("grpc server: serve: %v", err)
	case <-ctx.Done():
		grpcServer.GracefulStop()
		return errors.New("grpc server: context canceled")
	}
}
