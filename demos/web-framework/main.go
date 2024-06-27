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
	r.LoadHTMLGlob("templates/*")
	v1 := r.Group("")
	v1.Use(kit.Recovery())
	{
		v1.GET("/", func(c *kit.Context) {
			c.HTML(http.StatusOK, "css.tmpl", map[string]interface{}{
				"name": "jianjustin",
			})
		})
		// index out of range for testing Recovery()
		v1.GET("/panic", func(c *kit.Context) {
			names := []string{"geektutu"}
			c.String(http.StatusOK, names[100])
		})
		v1.Static("/assets", "./static")
	}

	r.Run(":9999")
}
