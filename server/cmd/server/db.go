package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
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
	err = initSpecifications(cfg.db)
	if err != nil {
		log.Fatalf("error initializing specifications: %s", err)
	}
	log.Println("specifications registered successfully")
}

func initSpecifications(db *database.Queries) error {
	for _, spec := range audit.SpecDefs {
		_, err := db.CreateSpecification(context.Background(), database.CreateSpecificationParams{
			Code:        spec.Code,
			Description: spec.Description,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "23505") {
				return fmt.Errorf("failed to create spec in database: %w", err)
			}
		}
	}
	return nil
}
