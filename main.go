package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Request path = %q\n", request.URL.Path)
	})

	r.GET("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "ping")
	})

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
