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
	"go.guide/div-grpc-service/middleware"
	"go.guide/div-grpc-service/pb"
	"go.guide/div-grpc-service/service"
	"go.guide/div-grpc-service/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	http1 "net/http"
	"os"
)

var (
	port    = flag.Int("port", 50054, "The server port")
	restful = flag.Int("restful", 8084, "the port to restful serve on")
)

func main() {
	flag.Parse()
	logger := middleware.InitLogger(os.Stdout)
	g, ctx := errgroup.WithContext(context.Background())

	svc := service.NewDivService()
	svc = service.LoggingDivServiceMiddleware(logger)(svc)
	svc = EtcdRegisterDivServiceMiddleware(logger)(svc)
	endpoints := transport.Endpoints(svc)

	level.Info(logger).Log("msg", fmt.Sprintf("grpc server start at 127.0.0.1:%v", *port))

	g.Go(func() error {
		return ServeGRPC(ctx, &endpoints, fmt.Sprintf("127.0.0.1:%d", *port), log.With(logger, "transport", "GRPC"))
	})
	g.Go(func() error {
		return ServeHTTP(ctx, &endpoints, fmt.Sprintf("127.0.0.1:%d", *restful), log.With(logger, "transport", "HTTP"))
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

func EtcdRegisterDivServiceMiddleware(logger log.Logger) service.DivServiceMiddleware {
	return func(next service.DivService) service.DivService {
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
