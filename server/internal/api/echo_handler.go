package api

import (
	"io"
	"net/http"
)

func HandleEcho(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error decoding parameters", &err)
		return
	}

	respondWithJSON(w, http.StatusCreated, string(body))
}
