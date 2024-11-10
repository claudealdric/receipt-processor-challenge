package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	http.ServeMux
}

func NewRouter(server *Server) *Router {
	router := Router{}
	router.Get("/{$}", func(w http.ResponseWriter, r *http.Request) {})
	router.Get("/receipts/{id}/points", server.HandleGetPoints)
	return &router
}

func (r *Router) Get(pattern string, handlerFunc http.HandlerFunc) {
	r.getHandlerFuncPattern(http.MethodGet, pattern, handlerFunc)
}

func (r *Router) getHandlerFuncPattern(
	method, pattern string,
	handlerFunc http.HandlerFunc,
) {
	r.HandleFunc(fmt.Sprintf("%s %s", method, pattern), handlerFunc)
}
