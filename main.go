package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Just1a2Noob/accouting-journal-app/internal/database"
	"github.com/Just1a2Noob/accouting-journal-app/internals/api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	fmt.Printf("DB_URL: %s\n", os.Getenv("DB_URL"))
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database : %s", err)
	}

	dbQueries := database.New(db)

	secret := os.Getenv("SECRET")
	apicfg := api.ApiConfig{
		Database: *dbQueries,
		Secret:   secret,
	}

	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/users", apicfg.HandlerCreateUser)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
