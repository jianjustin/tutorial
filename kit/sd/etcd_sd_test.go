package sd_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	"io"
	"os"
	"testing"
	"time"
)

func TestForNewEtcdClient(t *testing.T) {
	client, err := etcdv3.NewClient(
		context.Background(),
		[]string{"http://127.0.0.1:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}
	if client == nil {
		t.Fatal("expected new Client, got nil")
	}
	err = client.Register(etcdv3.Service{Key: "key1", Value: "value1"})
	if err != nil {
		return
	}
}

func TestForGetInstancer(t *testing.T) {
	prefix := "/services/foosvc/"
	instance := "127.0.0.1:8080"
	instance1 := "127.0.0.1:8081"
	client, err := etcdv3.NewClient(
		context.Background(),
		[]string{"http://127.0.0.1:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}
	if client == nil {
		t.Fatal("expected new Client, got nil")
	}
	client.Register(etcdv3.Service{Key: prefix + instance, Value: instance})
	client.Register(etcdv3.Service{Key: prefix + instance1, Value: instance1})

	logger := log.NewLogfmtLogger(os.Stderr)
	//根据service前缀获取一组实例
	instancer, err := etcdv3.NewInstancer(client, "/services/foosvc", logger)
	endpointer := sd.NewEndpointer(instancer, barFactory, logger)
	//将一组实例增加负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	marshal, err := json.Marshal(endpointer)
	logger.Log("msg", "get instancer", "endpointer", marshal)

	retry := lb.Retry(3, 3*time.Second, balancer)
	req := struct{}{}
	if _, err = retry(context.Background(), req); err != nil {
		logger.Log("retry err", err)
	}
}

func barFactory(s string) (endpoint.Endpoint, io.Closer, error) {
	fmt.Println("barFactory：" + s)
	return endpoint.Nop, nil, nil
}
