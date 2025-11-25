package handlers

import "gintama/internal/app"

type Handlers struct {
	Health  healthHandler
	Role    roleHandler
	User    userHandler
	Auth    authHandler
	Session sessionHandler
}

func New(app *app.Application) Handlers {
	return Handlers{
		Health:  healthHandler{app: app},
		Role:    roleHandler{app: app},
		User:    userHandler{app: app},
		Auth:    authHandler{app: app},
		Session: sessionHandler{app: app},
	}
}
