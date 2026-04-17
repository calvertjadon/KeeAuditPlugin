package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err *error) {
	if err != nil {
		msg = fmt.Sprintf("%s: %s", msg, *err)
	}
	log.Println(msg)
	type returnVals struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, returnVals{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}

func unmarshalParams[T any](requestBody io.ReadCloser, returnVals *T) error {
	decoder := json.NewDecoder(requestBody)
	defer requestBody.Close()

	return decoder.Decode(returnVals)
}
