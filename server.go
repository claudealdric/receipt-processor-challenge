package main

import "net/http"

type Server struct {
	http.Handler
}

func NewServer(router *Router) *Server {
	server := Server{}
	server.Handler = router
	return &server
}
