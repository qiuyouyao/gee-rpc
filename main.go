package main

import (
	"fmt"
	"gee"
	"net/http"
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

	adminGroup := r.Group("/admin")
	adminGroup.GET("/a", func(context *gee.Context) {
		_, _ = fmt.Fprintf(context.Writer, "Request path = %q\n", context.Request.URL.Path)
	})

	adminGroup.GET("/b/:id", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{
			"message": "success",
			"id":      context.Params["id"],
		})
	})

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
