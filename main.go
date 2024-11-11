package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudealdric/receipt-processor-challenge/api"
	"github.com/claudealdric/receipt-processor-challenge/data"
)

const port = 8080

func main() {
	store := data.NewInMemoryStore()
	server := api.NewServer(store)
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)
	log.Fatal(err)
}
