package main

import (
	"Blogger/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Port seems to be empty or unset")
	}

	dbURL := os.Getenv("CONN")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/users", apiCfg.handleUsersCreate)
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
