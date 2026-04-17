package api

import (
	"io"
	"net/http"
)

func (a *App) echo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (a *App) runAudit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type parameters struct {
		codes []string
	}

	params := parameters{}
	err := unmarshalParams(r.Body, &params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error decoding parameters", &err)
		return
	}

	auditID, err := a.auditor.Trigger(r.Context(), params.codes)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create audit", &err)
		return
	}

	respondWithJSON(w, http.StatusCreated, auditDto{ID: auditID})
}
