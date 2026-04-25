package mq

import (
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp.Channel
}

func NewPublisher(client *Client) (*Publisher, error) {
	ch, err := client.Channel()
	if err != nil {
		return nil, err
	}
	return &Publisher{channel: ch}, nil
}

func (p *Publisher) PublishStartAudit(cmd *audit.StartAuditCommand) error {
	return publishJSON(
		p.channel,
		"amq.direct",
		"audit.run",
		cmd,
	)
}
