package mq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
)

type ResultsConsumer struct {
	Exchange  string
	QueueName string
	Key       string
}

func (rc *ResultsConsumer) handle(results []audit.AuditResult) AckType {
	fmt.Println(results)
	return Ack
}

func (rc *ResultsConsumer) unmarshalResults(data []byte) ([]audit.AuditResult, error) {
	var results []audit.AuditResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (rc *ResultsConsumer) Subscribe(client *Client) {
	err := subscribe(
		client.Connection,
		// "amq.topic",
		// "audit.results",
		// "audit.results.*",
		rc.Exchange,
		rc.QueueName,
		rc.Key,
		SimpleQueueDurable,
		rc.handle,
		rc.unmarshalResults,
	)
	log.Println(err)
}
