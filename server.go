package main

import "net/http"

type Server struct {
	store Store
	http.Handler
}

func NewServer(store Store) *Server {
	server := Server{store: store}
	router := NewRouter(&server)
	server.Handler = router
	return &server
}
