package audit

import "github.com/google/uuid"

type Audit struct {
	ID    uuid.UUID
	Codes []string
}
