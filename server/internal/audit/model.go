package audit

import "github.com/google/uuid"

type Specification struct {
	Code string `json:"code"`
}

type ComplexityRequirement struct {
	Specification
	Threshold int32 `json:"threshold"`
}

type Audit struct {
	ID           uuid.UUID
	Requirements []ComplexityRequirement
}

type AuditResult struct {
	AuditID                  uuid.UUID `json:"audit_id"`
	EntryID                  uuid.UUID `json:"entry_id"`
	EntryTitle               string    `json:"entry_title"`
	GroupPath                string    `json:"group_path"`
	FailedSpecificationCodes []string  `json:"failed_specs"`
}
