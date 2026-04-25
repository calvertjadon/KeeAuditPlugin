package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/api"
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/config"
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/database"
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/mq"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("the server is starting up")

	godotenv.Load()
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DB.URL)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}

	client, err := mq.NewClient(cfg.MQ.URL)
	if err != nil {
		log.Fatalf("error connecting to rabbitmq: %s", err)
	}

	p, err := mq.NewPublisher(client)
	if err != nil {
		log.Fatalf("error creating pub channel: %s", err)
	}

	c := mq.ResultsConsumer{
		Exchange:  "amq.topic",
		QueueName: "audit.results",
		Key:       "audit.results.*",
	}
	c.Subscribe(client)

	auditRepo := database.NewAuditRepo(db)

	auditUseCase := audit.NewUseCase(auditRepo)
	auditHandler := api.NewAuditHandler(auditUseCase, p)
	router := api.NewRouter(auditHandler)

	log.Fatal(http.ListenAndServe(":"+cfg.App.Port, router))
}
