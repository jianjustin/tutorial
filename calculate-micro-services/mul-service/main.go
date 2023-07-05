package main

import (
	"github.com/go-kit/kit/sd/etcdv3"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"go.guide/mul-service/endpoint"
	"go.guide/mul-service/model"
	"go.guide/mul-service/register"
	"go.guide/mul-service/service"

	"net/http"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	svc := service.BaseMulService{}

	mulHandler := httptransport.NewServer(
		endpoint.MakeMulEndpoint(svc),
		model.DecodeMulRequest,
		model.EncodeResponse,
	)

	mulAfterAddHandler := httptransport.NewServer(
		endpoint.MakeMulAfterAddEndpoint(svc),
		model.DecodeMulRequest,
		model.EncodeResponse,
	)

	http.Handle("/mul", mulHandler)
	http.Handle("/mulAfterAdd", mulAfterAddHandler)
	logger.Log("Server started on port 8081")
	r := register.GetEtcdRegister()
	if r == nil {
		logger.Log("get register client failed")
		return
	}
	err := r.Register(etcdv3.Service{Key: "/services/mul/", Value: ":8081"})
	if err != nil {
		logger.Log("register service failed")
		return
	}
	defer r.Deregister(etcdv3.Service{Key: "/services/mul/", Value: ":8081"})

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		logger.Log("ListenAndServe: ", err)
	}
}
