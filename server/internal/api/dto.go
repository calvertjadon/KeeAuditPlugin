package api

import (
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
	"github.com/google/uuid"
)

type createAuditParams struct {
	Requirements []struct {
		SpecificationCode string `json:"code"`
		Threshold         int32  `json:"threshold"`
	} `json:"requirements"`
}

func (dto *createAuditParams) ToComplexityRequirement() []audit.ComplexityRequirement {
	var results []audit.ComplexityRequirement

	for _, r := range dto.Requirements {
		spec := audit.Specification{
			Code: r.SpecificationCode,
		}
		req := audit.ComplexityRequirement{
			Specification: spec,
			Threshold:     r.Threshold,
		}
		results = append(results, req)
	}

	return results
}

type createAuditResponse struct {
	ID uuid.UUID `json:"audit_id"`
}
