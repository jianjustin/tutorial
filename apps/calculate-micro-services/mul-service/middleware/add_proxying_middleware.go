package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/sony/gobreaker"
	"go.guide/mul-service/model"
	"go.guide/mul-service/service"
	"golang.org/x/time/rate"
	"net/url"
	"strings"
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
		qps         = 100                    // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
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

	// Now, build a single, retrying, load-balancing endpoint out of all of
	// those individual endpoints.
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	// And finally, return the ServiceMiddleware, implemented by proxymw.
	return func(next service.AddService) service.AddService {
		return AddServiceProxy{ctx, retry}
	}
}

func makeAddServiceProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/add"
	}
	return httptransport.NewClient(
		"GET",
		u,
		model.EncodeRequest,
		model.DecodeAddResponse,
	).Endpoint()
}

type AddServiceProxy struct {
	ctx context.Context
	E   endpoint.Endpoint
}

// Add 代理AddService的Add方法
func (proxy AddServiceProxy) Add(a int) (int, error) {
	response, err := proxy.E(proxy.ctx, model.AddRequest{A: a})
	if err != nil {
		return 0, err
	}

	str, _ := json.Marshal(response)

	resp := &model.AddResponse{}
	err = json.Unmarshal(str, resp)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}
