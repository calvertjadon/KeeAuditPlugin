package audit

import (
	"context"
	"fmt"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/database"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Auditor struct {
	conn *amqp.Connection
	db   *database.Queries
}

func NewAuditor(db *database.Queries, conn *amqp.Connection) *Auditor {
	return &Auditor{conn: conn, db: db}
}

type auditDto struct {
	ID uuid.UUID `json:"id"`
}

func (a *Auditor) Trigger(ctx context.Context, codes []string) (uuid.UUID, error) {
	_, err := a.db.GetSpecificationsByCodes(ctx, codes)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid specifications codes: %w", err)
	}

	auditID, err := a.db.CreateAudit(context.Background())
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create audit: %w", err)
	}

	if err := a.db.AttachAuditSpecs(ctx, database.AttachAuditSpecsParams{
		AuditID: auditID,
		Column2: codes,
	}); err != nil {
		return uuid.Nil, err
	}

	return auditID, nil
}
