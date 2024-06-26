package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/core-service/middleware"
	"go.guide/core-service/pb"
	"go.guide/core-service/service"
	"go.guide/core-service/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	port = flag.Int("port", 50050, "The server port")
)

func main() {
	flag.Parse()
	logger := middleware.InitLogger(os.Stdout)
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return InterruptHandler(ctx)
	})

	svc := service.NewCoreService()
	svc = service.LoggingCoreServiceMiddleware(logger)(svc)
	svc = EtcdRegisterAddServiceMiddleware(logger)(svc)
	endpoints := transport.Endpoints(svc)

	level.Info(logger).Log("msg", fmt.Sprintf("grpc server start at 127.0.0.1:%v", *port))

	g.Go(func() error {
		return ServeGRPC(ctx, &endpoints, fmt.Sprintf("127.0.0.1:%d", *port), log.With(logger, "transport", "GRPC"))
	})

	if err := g.Wait(); err != nil {
		log.With(logger, "level", "error").Log("error", err)
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
	pb.RegisterCoreServiceServer(grpcServer, server)
	log.With(logger, "level", "info").Log("listen on", addr)

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

func EtcdRegisterAddServiceMiddleware(logger log.Logger) service.CoreServiceMiddleware {
	return func(next service.CoreService) service.CoreService {
		r := middleware.GetEtcdRegister()
		if r == nil {
			level.Error(logger).Log("msg", "get register client failed")
			return next
		}
		err := r.Register(etcdv3.Service{Key: middleware.ServiceKey, Value: fmt.Sprintf("127.0.0.1:%d", *port)})
		if err != nil {
			level.Error(logger).Log("msg", "register service failed")
			return next
		}

		return next
	}
}
