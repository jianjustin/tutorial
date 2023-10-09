package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	log2 "jianjustin/add-grpc-service/middleware/log"
	register2 "jianjustin/add-grpc-service/middleware/register"
	pb2 "jianjustin/add-grpc-service/pb"
	"jianjustin/add-grpc-service/service"
	"jianjustin/add-grpc-service/transport"
	"jianjustin/add-grpc-service/transport/grpc"
	transporthttp "jianjustin/add-grpc-service/transport/http"
	"net"
	http1 "net/http"
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
	svc = log2.LoggingAddServiceMiddleware(logger)(svc)
	svc = register2.EtcdRegisterAddServiceMiddleware(register2.GetEtcdRegister(), logger)(svc)
	endpoints := transport.Endpoints(svc)

	log.With(logger, "level", "info").Log("msg", fmt.Sprintf("grpc server start at %s", register2.HostPort))

	g.Go(func() error {
		return ServeGRPC(ctx, &endpoints, register2.HostPort, log.With(logger, "transport", "GRPC"))
	})

	g.Go(func() error {
		return ServeHTTP(ctx, &endpoints, "localhost:18080", log.With(logger, "transport", "HTTP"))
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

// ServeGRPC serves gRPC requests.
func ServeGRPC(ctx context.Context, endpoints *transport.EndpointsSet, addr string, logger log.Logger) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server := transportgrpc.NewGRPCServer(endpoints,
		logger,
		opentracinggo.NoopTracer{},
	)
	grpcServer := grpc.NewServer()
	pb2.RegisterAddServiceServer(grpcServer, server)
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

func ServeHTTP(ctx context.Context, endpoints *transport.EndpointsSet, addr string, logger log.Logger) error {
	handler := transporthttp.NewHTTPHandler(endpoints,
		logger,
		opentracinggo.NoopTracer{},
	)
	httpServer := &http1.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.With(logger, "level", "info").Log("http listen on", addr)
	ch := make(chan error)
	go func() {
		ch <- httpServer.ListenAndServe()
	}()
	select {
	case err := <-ch:
		if err == http1.ErrServerClosed {
			return nil
		}
		return fmt.Errorf("http server: serve: %v", err)
	case <-ctx.Done():
		return httpServer.Shutdown(context.Background())
	}
}
