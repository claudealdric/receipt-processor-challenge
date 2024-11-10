package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	router := NewRouter()
	router.Get("/{$}", func(w http.ResponseWriter, r *http.Request) {})
	server := NewServer(router)
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)
	log.Fatal(err)
}
