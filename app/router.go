package app

import (
	"github.com/gorilla/handlers"
	"go-rest-api-template/app/handler"
)

func (a *App) configureRouter() {

	// Middlewares ...
	a.Router.Use(a.setRequestID)
	a.Router.Use(a.logRequest)
	a.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// User endpoints ...
	a.Post("/users", a.handleRequest(handler.UserCreate))
	a.Get("/users/{email}", a.handleRequest(handler.GetUserByEmail))
	a.Post("/sessions", a.handleRequest(handler.UserSessionsCreate))

	// Private endpoints ...
	private := a.Router.PathPrefix("/private").Subrouter()
	private.Use(a.authenticateUser)
	private.HandleFunc("/whoami", a.handleRequest(handler.Whoami))
}
