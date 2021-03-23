package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (a *App) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := a.Logger.WithFields(logrus.Fields{})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (a *App) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := a.SessionStore.Get(r, "user_session")
		if err != nil {
			a.error(w, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			a.error(w, http.StatusUnauthorized, errors.New("not authenticated"))
			return
		}

		u, err := a.Store.User().Find(id.(int))
		if err != nil {
			a.error(w, http.StatusUnauthorized, errors.New("not authenticated"))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "ctxKeyUser", u)))
	})
}

func (a *App) error(w http.ResponseWriter, code int, err error) {
	a.respond(w, code, map[string]string{"error": err.Error()})
}

func (a *App) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
