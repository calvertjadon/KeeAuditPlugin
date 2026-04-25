package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	*amqp.Connection
}

func New(url string) (*Client, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Client{Connection: conn}, nil
}
