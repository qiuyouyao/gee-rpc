package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/api/user/ping", func(ctx *gee.Context) {
		_, _ = fmt.Fprintf(ctx.Writer, "Request path = %q\n", ctx.Request.URL.Path)
	})

	r.GET("/ping/:id/:type", func(ctx *gee.Context) {
		_, _ = fmt.Fprint(ctx.Writer, gee.H{
			"message": "pong",
			"id":      ctx.Params["id"],
			"type":    ctx.Params["type"],
		})
	})

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
