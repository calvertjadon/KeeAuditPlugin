package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/database"
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/pubsub"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbit struct {
	host string
	port int
	user string
	pass string
}

func (r *rabbit) Connect() *amqp.Connection {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d", r.user, r.pass, r.host, r.port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatal("could not connect to RabbitMQ")
	}
	return conn
}

type apiConfig struct {
	db *database.Queries
}

func (cfg *apiConfig) initDatabase() {
	dbURL := os.Getenv("KEE_AUDIT_DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	cfg.db = database.New(db)

	specDefs := []struct {
		Code        string
		Description string
	}{
		{
			Code:        "entropy.min",
			Description: "specifies the minimum allowed entropy",
		},
		{
			Code:        "duplicates.none",
			Description: "specifies that a given password may only appear once in a given database",
		},
	}

	log.Println("registering specifications in database...")
	for _, spec := range specDefs {
		_, err := cfg.db.GetSpecificationByCode(context.Background(), spec.Code)
		if err != nil {
			log.Printf("%s spec already exists in database", spec.Code)
		}
	}
	log.Println("specifications registered successfully")
}

func main() {
	log.Println("the server is starting up")
	rabbitCfg := rabbit{
		host: "localhost",
		port: 5672,
		user: "guest",
		pass: "guest",
	}

	cfg := &apiConfig{}
	cfg.initDatabase()

	conn := rabbitCfg.Connect()
	defer conn.Close()
	defer log.Println("the server is shutting down")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	if _, _, err := pubsub.DeclareAndBind(
		conn,
		"amq.direct",
		"audit.results",
		"audit.results",
		pubsub.SimpleQueueDurable,
	); err != nil {
		log.Fatal(err)
	}

	if err := pubsub.SubscribeJSON(
		conn,
		"amq.direct",
		"audit.results",
		"audit.results",
		pubsub.SimpleQueueDurable,
		handleAuditResults(),
	); err != nil {
		log.Fatal(err)
	}

	enterRepl(ch)
}
