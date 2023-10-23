package main

import (
	"fmt"
	"github.com/jianjustin/web-kit/kit"
	"net/http"
)

func main() {
	r := kit.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Println(w, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":9999")
	if err != nil {
		fmt.Println("Run error")
	}
}
