package main

import (
	"fmt"
	"github.com/jianjustin/web-kit/kit"
)

func main() {
	r := kit.New()
	r.GET("/", func(c *kit.Context) {
		fmt.Println(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
	})

	r.GET("/hello", func(c *kit.Context) {
		for k, v := range c.Req.Header {
			fmt.Println(c.Writer, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":9999")
	if err != nil {
		fmt.Println("Run error")
	}
}
