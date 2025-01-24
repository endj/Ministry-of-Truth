package main

import (
	"app/db"
	"app/handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define command-line flags
	schemaFile := flag.String("schema", "schema.sql", "Path to the schema file (default: schema.sql)")
	dbFile := flag.String("db", "app.db", "Path to the SQLite database file (default: app.db)")

	// Parse flags
	flag.Parse()

	// Print defaults if used
	fmt.Println("Using schema file:", *schemaFile)
	fmt.Println("Using database file:", *dbFile)

	// Initialize the database with the provided schema and database file paths
	DB := db.InitializeDB(*dbFile, *schemaFile)

	repo := db.UserRepo{
		DB: DB,
	}

	// Define RESTful endpoints
	http.HandleFunc("/profiles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProfilesHandler(w, r, repo)
		case http.MethodPost:
			handlers.CreateProfileHandler(w, r, repo)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
