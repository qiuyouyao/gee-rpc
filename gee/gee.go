package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request)

type Engine struct {
	routers map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{routers: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	routeKey := buildRouteKey(method, pattern)
	engine.routers[routeKey] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute("PUT", pattern, handler)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute("DELETE", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	routeKey := buildRouteKey(request.Method, request.URL.Path)
	if handler, ok := engine.routers[routeKey]; ok {
		handler(writer, request)
	} else {
		_, _ = fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	}
}

func buildRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}
