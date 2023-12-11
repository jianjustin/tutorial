package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jianjustin/frontend/proto/add"
	pb "github.com/jianjustin/frontend/proto/mul"
	"github.com/jianjustin/frontend/proto/sub"
	"net/http"
	"strconv"
)

func (fe *frontendServer) AddHandler(w http.ResponseWriter, r *http.Request) {
	a, b := int64(0), int64(0)
	a, _ = strconv.ParseInt(mux.Vars(r)["a"], 10, 64)
	b, _ = strconv.ParseInt(mux.Vars(r)["b"], 10, 64)
	result, _ := fe.addService.Add(r.Context(), &add.AddRequest{A: a, B: b})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) MulHandler(w http.ResponseWriter, r *http.Request) {
	a, b := int64(0), int64(0)
	a, _ = strconv.ParseInt(mux.Vars(r)["a"], 10, 64)
	b, _ = strconv.ParseInt(mux.Vars(r)["b"], 10, 64)
	result, _ := fe.mulService.Mul(r.Context(), &pb.MulRequest{A: a, B: b})

	fmt.Fprint(w, result.GetResult())
}

func (fe *frontendServer) SubHandler(w http.ResponseWriter, r *http.Request) {
	a, b := int64(0), int64(0)
	a, _ = strconv.ParseInt(mux.Vars(r)["a"], 10, 64)
	b, _ = strconv.ParseInt(mux.Vars(r)["b"], 10, 64)
	result, _ := fe.subService.Sub(r.Context(), &sub.SubRequest{A: a, B: b})

	fmt.Fprint(w, result.GetResult())
}
