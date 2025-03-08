package main

import (
	"app/db"
	"app/handlers"
	"flag"
	"log"
	"net/http"
)

type App struct {
	userRepo db.UserRepo
	postRepo db.PostRepo
}

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
}

func (a *App) profilesHandler(w http.ResponseWriter, r *http.Request) {
	setCommonHeaders(w)
	switch r.Method {
	case http.MethodGet:
		handlers.GetProfilesHandler(w, r, a.userRepo)
	case http.MethodPost:
		handlers.CreateProfileHandler(w, r, a.userRepo)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) postsHandler(w http.ResponseWriter, r *http.Request) {
	setCommonHeaders(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		handlers.GetPostsHandler(w, r, a.postRepo)
	case http.MethodPost:
		handlers.CreatePostHandler(w, r, a.postRepo)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	schemaFile := flag.String("schema", "schema.sql", "Path to the schema file (default: schema.sql)")
	dbFile := flag.String("db", "app.db", "Path to the SQLite database file (default: app.db)")
	flag.Parse()

	log.Printf("Using schema file : %s, database file: %s", *schemaFile, *dbFile)

	DB := db.InitializeDB(*dbFile, *schemaFile)
	app := &App{userRepo: db.UserRepo{DB: DB}, postRepo: db.PostRepo{DB: DB}}

	http.HandleFunc("/profiles", app.profilesHandler)
	http.HandleFunc("/posts", app.postsHandler)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
