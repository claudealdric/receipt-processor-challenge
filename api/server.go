package api

import (
	"net/http"

	"github.com/claudealdric/receipt-processor-challenge/data"
)

type Server struct {
	store data.Store
	http.Handler
}

func NewServer(store data.Store) *Server {
	server := Server{store: store}
	router := NewRouter(&server)
	server.Handler = router
	return &server
}
