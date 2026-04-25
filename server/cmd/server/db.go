package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/database"
)

func (cfg *apiConfig) initDatabase() {
	log.Println("connecting to database...")
	dbURL := os.Getenv("KEE_AUDIT_DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}
	cfg.db = database.New(db)
	log.Println("connected to database successfully")

	log.Println("registering specifications in database...")
	if err != nil {
		log.Fatalf("error initializing specifications: %s", err)
	}
	log.Println("specifications registered successfully")
}
