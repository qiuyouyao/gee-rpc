package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (router *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s ", method, pattern)
	routeKey := buildRouteKey(method, pattern)
	router.handlers[routeKey] = handler
}

func (router *router) handle(context *Context) {
	routeKey := buildRouteKey(context.Method, context.Path)
	if handler, ok := router.handlers[routeKey]; ok {
		handler(context)
	} else {
		_, err := context.String(http.StatusNotFound, "404 NOT FOUND: %s\n", context.Path)
		if err != nil {
			return
		}
	}
}

func buildRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}
