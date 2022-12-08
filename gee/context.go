package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Method     string
	Path       string
	StatusCode int
}

func newContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer:  writer,
		Request: request,
		Method:  request.Method,
		Path:    request.URL.Path,
	}
}

func (context *Context) PostForm(key string) string {
	return context.Request.FormValue(key)
}

func (context *Context) Query(key string) string {
	return context.Request.URL.Query().Get(key)
}

func (context *Context) Status(code int) {
	context.StatusCode = code
	context.Writer.WriteHeader(code)
}

func (context *Context) SetHeader(key string, value string) {
	context.Writer.Header().Set(key, value)
}

func (context *Context) String(statusCode int, format string, values ...interface{}) (n int, err error) {
	context.SetHeader("Content-type", "text/plain")
	context.Status(statusCode)
	return context.Writer.Write([]byte(fmt.Sprintf(format, values)))
}

func (context *Context) JSON(statusCode int, obj interface{}) {
	context.SetHeader("Content-type", "application/json")
	context.Status(statusCode)
	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), 500)
	}
}

func (context *Context) Data(statusCode int, data []byte) (n int, err error) {
	context.Status(statusCode)
	return context.Writer.Write(data)
}

func (context *Context) HTML(statusCode int, html string) (n int, err error) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(statusCode)
	return context.Writer.Write([]byte(html))
}
