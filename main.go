package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Port seems to be empty or unset")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", readinessHandler)
	mux.HandleFunc("GET /v1/err", errorHandler)

	server := http.Server{
		Addr:        fmt.Sprintf(":%s", PORT),
		Handler:     mux,
		ReadTimeout: 10 * time.Second,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
