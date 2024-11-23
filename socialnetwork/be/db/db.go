package db

import (
	"database/sql"
	"log"
)

func InitializeDB(dataSourceName string) *sql.DB {
	var err error
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	return db
}
