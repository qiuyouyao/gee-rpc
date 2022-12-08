package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(ctx *gee.Context) {
		_, _ = fmt.Fprintf(ctx.Writer, "Request path = %q\n", ctx.Request.URL.Path)
	})

	r.GET("/ping", func(ctx *gee.Context) {
		_, _ = fmt.Fprint(ctx.Writer, "pong")
	})

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
