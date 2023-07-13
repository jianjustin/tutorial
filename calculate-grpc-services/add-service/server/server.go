package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	log2 "go.guide/add-grpc-service/log"
	"go.guide/add-grpc-service/middleware"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/register"
	"go.guide/add-grpc-service/service"
	"go.guide/add-grpc-service/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.With(log2.InitLogger(os.Stdout))
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return InterruptHandler(ctx)
	})

	svc := service.NewAddService()
	svc = middleware.LoggingAddServiceMiddleware(logger)(svc)
	svc = middleware.EtcdRegisterAddServiceMiddleware(register.GetEtcdRegister(), logger)(svc)
	endpoints := transport.Endpoints(service.NewAddService())

	log.With(logger, "level", "info").Log("msg", fmt.Sprintf("grpc server start at %s", middleware.HostPort))

	g.Go(func() error {
		return ServeGRPC(ctx, &endpoints, middleware.HostPort, log.With(logger, "transport", "GRPC"))
	})

	if err := g.Wait(); err != nil {
		log.With(logger, "level", "error").Log("error", err)
	}
}

// InterruptHandler handles first SIGINT and SIGTERM and returns it as error.
func InterruptHandler(ctx context.Context) error {
	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-interruptHandler:
		return fmt.Errorf("signal received: %v", sig.String())
	case <-ctx.Done():
		return errors.New("signal listener: context canceled")
	}
}

func ServeGRPC(ctx context.Context, endpoints *transport.EndpointsSet, addr string, logger log.Logger) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// Here you can add middlewares for grpc server.
	server := transport.NewGRPCServer(endpoints,
		logger,
		opentracinggo.NoopTracer{},
	)
	grpcServer := grpc.NewServer()
	pb2.RegisterAddServiceServer(grpcServer, server)
	logger.Log("listen on", addr)
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
