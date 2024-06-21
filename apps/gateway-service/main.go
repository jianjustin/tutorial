package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/gateway-service/middleware"
	"go.guide/gateway-service/transport"
	"golang.org/x/sync/errgroup"
	http1 "net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	restful = flag.Int("restful", 8085, "the port to restful serve on")
)

func main() {
	flag.Parse()
	logger := middleware.InitLogger(os.Stdout)
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return InterruptHandler(ctx)
	})

	endpoints := transport.Endpoints()

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
