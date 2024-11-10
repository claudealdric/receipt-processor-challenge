package main

import "net/http"

type Server struct {
	http.Handler
}

func NewServer() *Server {
	server := Server{}
	router := NewRouter(&server)
	server.Handler = router
	return &server
}
