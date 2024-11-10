package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	store := InMemoryStore{}
	server := NewServer(&store)
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)
	log.Fatal(err)
}
