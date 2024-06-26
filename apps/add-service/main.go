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
	"go.guide/add-grpc-service/middleware"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
	"go.guide/add-grpc-service/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	http1 "net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	restful = flag.Int("restful", 8081, "the port to restful serve on")
)

func main() {
	flag.Parse()
	logger := middleware.InitLogger(os.Stdout)
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return InterruptHandler(ctx)
	})

	svc := service.NewAddService()
	svc = service.LoggingAddServiceMiddleware(logger)(svc)
	svc = EtcdRegisterAddServiceMiddleware(logger)(svc)
	endpoints := transport.Endpoints(svc)

	level.Info(logger).Log("msg", fmt.Sprintf("grpc server start at 127.0.0.1:%v", *port))

	g.Go(func() error {
		return ServeGRPC(ctx, &endpoints, fmt.Sprintf("127.0.0.1:%d", *port), log.With(logger, "transport", "GRPC"))
	})

	g.Go(func() error {
		return ServeHTTP(ctx, &endpoints, fmt.Sprintf("127.0.0.1:%d", *restful), log.With(logger, "transport", "HTTP"))
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

	server := transport.NewGRPCServer(endpoints,
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
	handler := transport.NewHTTPHandler(endpoints,
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
		if errors.Is(err, http1.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("http server: serve: %v", err)
	case <-ctx.Done():
		return httpServer.Shutdown(context.Background())
	}
}

func EtcdRegisterAddServiceMiddleware(logger log.Logger) service.AddServiceMiddleware {
	return func(next service.AddService) service.AddService {
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
