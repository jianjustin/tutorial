package middleware

import (
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"go.guide/add-grpc-service/register"
	"go.guide/add-grpc-service/service"
)

const HostPort string = "localhost:8002"
const ServiceKey string = "/services/add/"

func EtcdRegisterAddServiceMiddleware(e etcdv3.Client, logger log.Logger) service.AddServiceMiddleware {
	return func(next service.AddService) service.AddService {
		r := register.GetEtcdRegister()
		if r == nil {
			log.With(logger, "level", "error").Log("msg", "get register client failed")
			return next
		}
		err := r.Register(etcdv3.Service{Key: ServiceKey, Value: HostPort})
		if err != nil {
			log.With(logger, "level", "error").Log("msg", "register service failed")
			return next
		}
		defer r.Deregister(etcdv3.Service{Key: ServiceKey, Value: HostPort})

		return next
	}
}
