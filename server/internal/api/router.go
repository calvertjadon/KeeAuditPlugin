package api

func (a *App) configureRoutes() {
	a.mux.HandleFunc("POST /echo", a.echo)
	a.mux.HandleFunc("POST /run", a.runAudit)
}
