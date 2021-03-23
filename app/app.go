package app

import (
	"github.com/sirupsen/logrus"
	"go-rest-api-template/app/store"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type App struct {
	Router       *mux.Router
	Logger       *logrus.Logger
	Store        store.Store
	SessionStore sessions.Store
}

func (a *App) Initialize(store store.Store, sessionStore sessions.Store) *mux.Router {
	a.Router = mux.NewRouter()
	a.Logger = logrus.New()
	a.Store = store
	a.SessionStore = sessionStore

	a.configureRouter()

	return a.Router
}
