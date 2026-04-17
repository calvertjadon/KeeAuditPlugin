package main

import (
	"fmt"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/pubsub"
	"github.com/google/uuid"
)

type AuditResultDto struct {
	EntryID              uuid.UUID `json:"entry_id"`
	EntryTitle           string    `json:"entry_title"`
	GroupPath            string    `json:"group_path"`
	IsCompliant          bool      `json:"is_compliant"`
	FailedSpecifications []string  `json:"failed_specification"`
}

func handleAuditResults() func([]AuditResultDto) pubsub.AckType {
	return func(ars []AuditResultDto) pubsub.AckType {
		defer fmt.Println("> ")
		for _, ar := range ars {
			fmt.Println(ar)
		}
		return pubsub.Ack
	}
}
