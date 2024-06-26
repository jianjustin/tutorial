package proxying

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/sony/gobreaker"
	"go.guide/div-grpc-service/middleware"
	"go.guide/div-grpc-service/pb"
	"go.guide/div-grpc-service/service"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
)

func ProxyingMiddleware(ctx context.Context, serviceName string, logger log.Logger) service.CoreServiceMiddleware {
	// If instances is empty, don't proxy.
	if serviceName == "" {
		logger.Log("proxy_to", "none")
		return func(next service.CoreService) service.CoreService { return next }
	}

	var (
		qps = 100
	)

	var (
		endpointer sd.FixedEndpointer
	)

	client, err := etcdv3.NewClient(
		context.Background(),
		[]string{middleware.EtcdHost},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	if err != nil {
		logger.Log("unexpected error creating client: %v", err)
		return func(next service.CoreService) service.CoreService { return next }
	}
	if client == nil {
		logger.Log("expected new Client, got nil")
		return func(next service.CoreService) service.CoreService { return next }
	}
	instances, err := client.GetEntries(serviceName)
	if err != nil || len(instances) == 0 {
		return func(next service.CoreService) service.CoreService { return next }
	}

	logger.Log("proxy_to", fmt.Sprint(instances))
	for _, instance := range instances {
		var e endpoint.Endpoint
		e = makeCoreServiceProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	return func(next service.CoreService) service.CoreService {
		return CodeServiceProxy{context.Background(), endpointer[0]}
	}
}

func makeCoreServiceProxy(ctx context.Context, instance string) endpoint.Endpoint {
	// 连接到服务实例
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	cc, err := grpc.DialContext(ctx, instance, grpc.WithInsecure())
	if err != nil {
		fmt.Println("unable to Dial: %+v", err)
	}

	//创建代理端点
	return grpctransport.NewClient(
		cc,
		"pb.CoreService",
		"Random",
		_Encode_Grpc_Random_Request,
		_Decode_Grpc_Random_Response,
		&pb.RandomResponse{},
	).Endpoint()
}
