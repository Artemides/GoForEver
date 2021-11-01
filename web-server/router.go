package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, ePath := r.rules[path]
	handler, eMethod := r.rules[path][method]
	return handler, eMethod, ePath
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, eMethod, ePath := r.FindHandler(req.URL.Path, req.Method)
	if !ePath {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !eMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, req)

}
