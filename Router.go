package peiqi

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

// 用于存储路由
type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]HandlerFunc)
	}
	r.routes[path][method] = handler
}

/*
		ServeHTTP 实现了 :
		type Handler interface {
	    	ServeHTTP(w http.ResponseWriter, r *http.Request)
		}
*/
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	methodRoutes, exists := r.routes[req.URL.Path]
	if !exists {
		http.NotFound(w, req)
		return
	}
	handler, exists := methodRoutes[req.Method]
	if !exists {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	handler(context)
}
