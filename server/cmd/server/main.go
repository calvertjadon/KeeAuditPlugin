package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	log.Println("the server is starting up")

	godotenv.Load()

	r, err := newRabbitMQ(
		"localhost",
		5672,
		"guest",
		"guest",
	)
	if err != nil {
		log.Fatal(err)
	}
	r.configureQueues()

	cfg := &apiConfig{}
	cfg.initDatabase()

	defer r.Conn.Close()
	defer log.Println("the server is shutting down")

	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	enterRepl(ch)
}
