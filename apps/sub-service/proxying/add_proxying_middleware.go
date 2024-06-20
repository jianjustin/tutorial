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
	"go.guide/add-grpc-service/pb"
	"go.guide/sub-grpc-service/service"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
)

func ProxyingMiddleware(ctx context.Context, serviceName string, logger log.Logger) service.AddServiceMiddleware {
	// If instances is empty, don't proxy.
	if serviceName == "" {
		logger.Log("proxy_to", "none")
		return func(next service.AddService) service.AddService { return next }
	}

	// Set some parameters for our client.
	var (
		qps = 100 // beyond which we will return an error
		//maxAttempts = 3                 // per request, before giving up
		//maxTime     = 250 * time.Second // wallclock time, before giving up
	)

	// Otherwise, construct an endpoint for each instance in the list, and add
	// it to a fixed set of endpoints. In a real service, rather than doing this
	// by hand, you'd probably use package sd's support for your service
	// discovery system.
	var (
		endpointer sd.FixedEndpointer
	)

	client, err := etcdv3.NewClient(
		context.Background(),
		[]string{"http://127.0.0.1:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	if err != nil {
		logger.Log("unexpected error creating client: %v", err)
		return func(next service.AddService) service.AddService { return next }
	}
	if client == nil {
		logger.Log("expected new Client, got nil")
		return func(next service.AddService) service.AddService { return next }
	}
	instances, err := client.GetEntries(serviceName)
	if err != nil || len(instances) == 0 {
		return func(next service.AddService) service.AddService { return next }
	}

	logger.Log("proxy_to", fmt.Sprint(instances))
	for _, instance := range instances {
		var e endpoint.Endpoint
		e = makeAddServiceProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	// And finally, return the ServiceMiddleware, implemented by proxymw.
	return func(next service.AddService) service.AddService {
		return AddServiceProxy{context.Background(), endpointer[0]}
	}
}

func makeAddServiceProxy(ctx context.Context, instance string) endpoint.Endpoint {
	// 连接到服务实例
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	cc, err := grpc.DialContext(ctx, instance, grpc.WithInsecure())
	if err != nil {
		fmt.Println("unable to Dial: %+v", err)
	}

	//创建代理端点
	return grpctransport.NewClient(
		cc,
		"pb.AddService",
		"Add",
		_Encode_Add_Request,
		_Decode_Add_Response,
		&pb.AddResponse{},
	).Endpoint()
}
