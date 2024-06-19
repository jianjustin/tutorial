package main

import (
	"github.com/go-kit/kit/sd/etcdv3"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.guide/add-service/endpoint"
	"go.guide/add-service/register"
	"go.guide/add-service/service"
	"log"
	"net/http"
)

func main() {
	svc := service.BaseAddService{}

	addHandler := httptransport.NewServer(
		endpoint.MakeAddEndpoint(svc),
		endpoint.DecodeAddRequest,
		endpoint.EncodeResponse,
	)

	addAfterMulHandler := httptransport.NewServer(
		endpoint.MakeAddAfterMulEndpoint(svc),
		endpoint.DecodeAddRequest,
		endpoint.EncodeResponse,
	)

	http.Handle("/add", addHandler)
	http.Handle("/addAfterMul", addAfterMulHandler)

	log.Println("Server started on port 8080")
	r := register.GetEtcdRegister()
	if r == nil {
		log.Println("get register client failed")
		return
	}
	err := r.Register(etcdv3.Service{Key: "/services/add/", Value: ":8080"})
	if err != nil {
		log.Println("register service failed")
		return
	}
	defer r.Deregister(etcdv3.Service{Key: "/services/add/", Value: ":8080"})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
