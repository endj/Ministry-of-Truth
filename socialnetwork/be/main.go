package main

import (
	"app/db"
	"app/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "user")
	dbPassword := getEnv("DB_PASSWORD", "user")
	dbName := getEnv("DB_NAME", "user")

	dataSourceConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db.InitializeDB(dataSourceConfig)

	http.HandleFunc("/profiles", handlers.HandleGetProfiles)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
