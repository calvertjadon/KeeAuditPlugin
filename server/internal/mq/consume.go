package mq

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	SimpleQueueTransient = iota
	SimpleQueueDurable
)

type AckType int

const (
	Ack AckType = iota
	NackRequeue
	NackDiscard
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	queue, err := ch.QueueDeclare(
		queueName,
		queueType == SimpleQueueDurable,
		queueType == SimpleQueueTransient,
		queueType == SimpleQueueTransient,
		false,
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = ch.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return ch, queue, nil
}

func subscribe[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
	handler func(T) AckType,
	unmarshaller func([]byte) (T, error),
) error {
	ch, queue, err := DeclareAndBind(
		conn,
		exchange,
		queueName,
		key,
		queueType,
	)
	if err != nil {
		return fmt.Errorf("could not declare and bind queue: %w", err)
	}

	if err := ch.Qos(10, 0, true); err != nil {
		return fmt.Errorf("could not set prefetch count: %w", err)
	}

	deliveries, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not consume messages: %w", err)
	}

	go func() {
		defer ch.Close()
		for delivery := range deliveries {
			target, err := unmarshaller(delivery.Body)
			if err != nil {
				fmt.Printf("could not unmarshal delivery: %v\n", err)
				continue
			}
			at := handler(target)

			switch at {
			case Ack:
				log.Printf("ack %s", key)
				delivery.Ack(false)
			case NackRequeue:
				log.Printf("nack requeue %s", key)
				delivery.Nack(false, true)
			case NackDiscard:
				log.Printf("nack discard %s", key)
				delivery.Nack(false, false)
			}

		}
	}()
	return nil
}

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
	handler func(T) AckType,
) error {
	return subscribe(
		conn,
		exchange,
		queueName,
		key,
		queueType,
		handler,
		func(data []byte) (T, error) {
			var target T
			log.Println(string(data))
			err := json.Unmarshal(data, &target)
			return target, err
		})
}
