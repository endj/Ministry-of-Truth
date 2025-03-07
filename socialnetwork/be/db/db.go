package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func InitializeDB(dataSourceName string, schemaFile string) *sql.DB {
	var err error
	log.Println("Opening sql db")
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatalf("Error testing database connection: %v", err)
	}
	applySchema(db, schemaFile)
	return db
}

func applySchema(db *sql.DB, schemaFile string) {
	log.Println("Applying schema from", schemaFile)
	data, err := os.ReadFile(schemaFile)
	if err != nil {
		log.Fatalf("Error reading schema file: %v", err)
	}

	schema := string(data)
	statements := strings.Split(schema, ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			log.Fatalf("Error executing statement: %v", err)
		}
	}
	log.Println("Schema applied successfully")
}
