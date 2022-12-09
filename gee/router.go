package gee

import (
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
	routeKey := buildRouteKey(method, pattern)
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)
	router.handlers[routeKey] = handler
}

func (router *router) handle(context *Context) {
	n, params := router.getRoute(context.Method, context.Path)
	if n != nil {
		context.Params = params
		routeKey := buildRouteKey(context.Method, n.pattern)
		router.handlers[routeKey](context)
	} else {
		_, err := context.String(http.StatusNotFound, "404 NOT FOUND: %s\n", context.Path)
		if err != nil {
			return
		}
	}
}

func (router *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			} else if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[:index], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
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
