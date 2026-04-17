package main

import (
	"fmt"
	"log"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbit struct {
	Conn *amqp.Connection
}

func newRabbitMQ(host string, port int, user string, pass string) (*rabbit, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d", user, pass, host, port)
	r := rabbit{}
	if err := r.Connect(connStr); err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *rabbit) Connect(connStr string) error {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return fmt.Errorf("could not connect to RabbitMQ: %w", err)
	}
	r.Conn = conn
	return nil
}

func (r *rabbit) configureQueues() {
	if _, _, err := pubsub.DeclareAndBind(
		r.Conn,
		"amq.direct",
		"audit.results",
		"audit.results",
		pubsub.SimpleQueueDurable,
	); err != nil {
		log.Fatal(err)
	}

	if err := pubsub.SubscribeJSON(
		r.Conn,
		"amq.direct",
		"audit.results",
		"audit.results",
		pubsub.SimpleQueueDurable,
		handleAuditResults(),
	); err != nil {
		log.Fatal(err)
	}
}
