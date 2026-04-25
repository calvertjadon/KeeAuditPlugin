package audit

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(context.Context, *Audit) error
}

type UseCase struct {
	repo Repository
}

func NewUseCase(r Repository) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) StartAudit(ctx context.Context, codes []string) (*Audit, error) {
	a := &Audit{
		ID:    uuid.New(),
		Codes: codes,
	}

	if err := uc.repo.Save(ctx, a); err != nil {
		return nil, err
	}

	return a, nil
}
