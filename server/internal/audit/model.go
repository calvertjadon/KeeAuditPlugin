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
