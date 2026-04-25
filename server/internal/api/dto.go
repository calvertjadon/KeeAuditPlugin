package api

import "github.com/google/uuid"

type createAuditParams struct {
	Codes []string `json:"codes"`
}
type createAuditResponse struct {
	ID uuid.UUID `json:"audit_id"`
}
