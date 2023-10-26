package main

import (
	"github.com/jianjustin/web-kit/kit"
	"log"
	"net/http"
	"time"
)

func onlyForV2() kit.HandlerFunc {
	return func(c *kit.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		//c.Status(http.StatusInternalServerError)
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := kit.New()
	//r.Use(kit.Logger()) // global midlleware
	r.GET("/", func(c *kit.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *kit.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
