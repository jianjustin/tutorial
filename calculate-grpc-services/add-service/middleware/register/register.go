package register

import (
	"context"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"go.guide/add-grpc-service/service"
	"time"
)

const HostPort string = "localhost:8002"
const ServiceKey string = "/services/add/"

func GetEtcdRegister() etcdv3.Client {
	client, _ := etcdv3.NewClient(
		context.Background(),
		[]string{"http://127.0.0.1:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	return client
}

func EtcdRegisterAddServiceMiddleware(e etcdv3.Client, logger log.Logger) service.AddServiceMiddleware {
	return func(next service.AddService) service.AddService {
		r := GetEtcdRegister()
		if r == nil {
			log.With(logger, "level", "error").Log("msg", "get register client failed")
			return next
		}
		err := r.Register(etcdv3.Service{Key: ServiceKey, Value: HostPort})
		if err != nil {
			log.With(logger, "level", "error").Log("msg", "register service failed")
			return next
		}
		//defer r.Deregister(etcdv3.Service{Key: ServiceKey, Value: HostPort})

		return next
	}
}
