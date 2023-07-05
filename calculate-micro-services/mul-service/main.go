package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"go.guide/mul-service/endpoint"
	"go.guide/mul-service/service"
	"log"
	"net/http"
)

func main() {
	svc := service.BaseStringService{}

	mulHandler := httptransport.NewServer(
		endpoint.MakeMulEndpoint(svc),
		endpoint.DecodeMulRequest,
		endpoint.EncodeResponse,
	)

	http.Handle("/mul", mulHandler)

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
