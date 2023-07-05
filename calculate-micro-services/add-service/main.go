package main

import (
	"go.guide/add-service/endpoint"
	"go.guide/add-service/service"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := service.BaseStringService{}

	addHandler := httptransport.NewServer(
		endpoint.MakeAddEndpoint(svc),
		endpoint.DecodeAddRequest,
		endpoint.EncodeResponse,
	)

	http.Handle("/add", addHandler)

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
