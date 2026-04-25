package audit

import "github.com/google/uuid"

type StartAuditCommand struct {
	AuditID      uuid.UUID               `json:"audit_id"`
	Requirements []ComplexityRequirement `json:"requirements"`
}

type CommandPublisher interface {
	PublishStartAudit(cmd *StartAuditCommand) error
}

func NewStartAuditCommand(a *Audit) *StartAuditCommand {
	return &StartAuditCommand{
		AuditID:      a.ID,
		Requirements: a.Requirements,
	}
}
