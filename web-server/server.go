package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}
func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, ePath := s.router.rules[path]
	if !ePath {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}
func (s *Server) AddMiddleWare(hf http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		hf = m(hf)
	}
	return hf
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
