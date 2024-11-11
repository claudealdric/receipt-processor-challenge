package api

import (
	"fmt"
	"net/http"
)

type Router struct {
	http.ServeMux
}

func NewRouter(server *Server) *Router {
	router := Router{}
	router.Get("/receipts/{id}/points", server.HandleGetPoints)
	router.Post("/receipts/process", server.HandleProcessReceipt)
	return &router
}

func (r *Router) Get(pattern string, handlerFunc http.HandlerFunc) {
	r.getHandlerFuncPattern(http.MethodGet, pattern, handlerFunc)
}

func (r *Router) Post(pattern string, handlerFunc http.HandlerFunc) {
	r.getHandlerFuncPattern(http.MethodPost, pattern, handlerFunc)
}

func (r *Router) getHandlerFuncPattern(
	method, pattern string,
	handlerFunc http.HandlerFunc,
) {
	r.HandleFunc(fmt.Sprintf("%s %s", method, pattern), handlerFunc)
}
