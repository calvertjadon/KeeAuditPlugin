package api

import (
	"net/http"

	"github.com/calvertjadon/KeeAuditPlugin/server/internal/audit"
	"github.com/calvertjadon/KeeAuditPlugin/server/internal/database"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	auditor *audit.Auditor
	mux     *http.ServeMux
}

func NewApp(db *database.Queries, conn *amqp.Connection) *App {
	a := &App{
		auditor: audit.NewAuditor(db, conn),
		mux:     http.NewServeMux(),
	}
	a.configureRoutes()
	return a
}

func (a *App) Run() {
	http.ListenAndServe(":8080", a.mux)
}
