package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
)

type AuditRepo struct {
	q *Queries
}

func NewAuditRepo(db *sql.DB) *AuditRepo {
	return &AuditRepo{
		q: New(db),
	}
}

func (r *AuditRepo) Save(ctx context.Context, a *audit.Audit) error {
	if _, err := r.q.GetSpecificationsByCodes(ctx, a.Codes); err != nil {
		return fmt.Errorf("invalid specifications codes: %w", err)
	}

	if _, err := r.q.CreateAudit(ctx, a.ID); err != nil {
		return err
	}

	if err := r.q.AttachAuditSpecs(ctx, AttachAuditSpecsParams{
		AuditID: a.ID,
		Column2: a.Codes,
	}); err != nil {
		return err
	}

	return nil
}
