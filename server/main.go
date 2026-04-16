package main

import (
	"fmt"
	"log"

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

func main() {
	log.Println("the server is starting up")
	rabbitCfg := rabbit{
		host: "localhost",
		port: 5672,
		user: "guest",
		pass: "guest",
	}

	conn := rabbitCfg.Connect()
	defer conn.Close()
	defer log.Println("the server is shutting down")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	enterRepl(ch)
}
