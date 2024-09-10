package main

import (
	"log"
	"net/http"

	"github.com/yourusername/lassie-retriever/internal/handler"
	"github.com/yourusername/lassie-retriever/internal/lassie"
)

func main() {
	lassieClient, err := lassie.NewClient("http://localhost:1234")
	if err != nil {
		log.Fatalf("Failed to create Lassie client: %v", err)
	}

	http.HandleFunc("/retrieve", handler.RetrieveHandler(lassieClient))

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
