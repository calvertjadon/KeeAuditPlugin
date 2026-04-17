package api

import "github.com/google/uuid"

type auditDto struct {
	ID uuid.UUID `json:"audit_id"`
}
