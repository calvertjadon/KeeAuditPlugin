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
	if _, err := r.q.CreateAudit(ctx, a.ID); err != nil {
		return err
	}

	for _, requirement := range a.Requirements {
		spec, err := r.q.GetSpecificationByCode(ctx, requirement.Code)
		if err != nil {
			return fmt.Errorf("invalid specification code: %w", err)
		}

		if err := r.q.AttachAuditSpecs(ctx, AttachAuditSpecsParams{
			AuditID:         a.ID,
			SpecificationID: spec.ID,
			Threshold:       requirement.Threshold,
		}); err != nil {
			return err
		}
	}

	return nil
}
