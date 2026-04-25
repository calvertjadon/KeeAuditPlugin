package api

import (
	"net/http"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
)

type auditHandler struct {
	uc *audit.UseCase
	p  audit.CommandPublisher
}

func NewAuditHandler(uc *audit.UseCase, p audit.CommandPublisher) *auditHandler {
	return &auditHandler{uc: uc, p: p}
}

func (h *auditHandler) HandleCreateAudit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := createAuditParams{}
	err := unmarshalParams(r.Body, &params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error decoding parameters", &err)
		return
	}

	a, err := h.uc.StartAudit(r.Context(), params.ToComplexityRequirement())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create audit", &err)
		return
	}

	cmd := audit.NewStartAuditCommand(a)

	if err := h.p.PublishStartAudit(cmd); err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to publish", &err)
		return
	}

	resp := createAuditResponse{ID: a.ID}
	respondWithJSON(w, http.StatusCreated, resp)
}
