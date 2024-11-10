package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	http.ServeMux
}

func NewRouter() *Router {
	return &Router{}
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
