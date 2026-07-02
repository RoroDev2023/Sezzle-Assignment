package main

import (
	"log"
	"net/http"
	"os"

	"sezzle-calculator/backend/internal/calculator"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := calculator.NewServer()
	addr := ":" + port

	log.Printf("calculator API listening on %s", addr)
	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Fatal(err)
	}
}

