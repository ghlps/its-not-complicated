package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type App struct {
	db *sql.DB
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello world!"))
}

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("it's not possible to open connection %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("something is wrong with the db: %v", err)
	}
	log.Printf("successfully connected with the db")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	log.Printf("listening on the port 8080")
	http.ListenAndServe(":8080", mux)
}
