package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (router *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	log.Printf("Route %4s - %s ", method, pattern)
	routeKey := buildRouteKey(method, pattern)
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)
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

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func buildRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}
