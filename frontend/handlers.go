package main

import (
	"encoding/json"
	"fmt"
	"github.com/jianjustin/frontend/proto/add"
	pb "github.com/jianjustin/frontend/proto/mul"
	"github.com/jianjustin/frontend/proto/sub"
	"net/http"
)

type requestBody struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

func (fe *frontendServer) AddHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, _ := fe.addService.Add(r.Context(), &add.AddRequest{A: body.A, B: body.B})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) MulHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, _ := fe.mulService.Mul(r.Context(), &pb.MulRequest{A: body.A, B: body.B})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) SubHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, _ := fe.subService.Sub(r.Context(), &sub.SubRequest{A: body.A, B: body.B})

	fmt.Fprint(w, result.GetResult())
}
