package api

import "net/http"

func NewRouter(audit *auditHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /audits", audit.HandleCreateAudit)
	mux.HandleFunc("GET /echo", HandleEcho)

	return mux
}
