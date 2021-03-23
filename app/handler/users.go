package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go-rest-api-template/app/model"
	"go-rest-api-template/app/store"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request, s store.Store, ss sessions.Store) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.User().Create(u); err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}

	u.Sanitize()
	respondJSON(w, http.StatusCreated, u)
}

func UserSessionsCreate(w http.ResponseWriter, r *http.Request, s store.Store, ss sessions.Store) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	u, err := s.User().FindByEmail(req.Email)
	if err != nil || !u.ComparePassword(req.Password) {
		respondError(w, http.StatusUnauthorized, errors.New("Incorrect email or password"))
		return
	}

	session, err := ss.Get(r, "user_session")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	session.Values["user_id"] = u.ID
	if err := ss.Save(r, w, session); err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func Whoami(w http.ResponseWriter, r *http.Request, s store.Store, ss sessions.Store) {
	respondJSON(w, http.StatusOK, model.CreateResponse(r.Context().Value("ctxKeyUser")))
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request, s store.Store, ss sessions.Store) {
	type request struct {
		Email string `json:"email"`
	}

	vars := mux.Vars(r)

	email := vars["email"]

	u, err := s.User().FindByEmail(email)
	if err != nil {
		respondError(w, http.StatusUnauthorized, errors.New("Incorrect email or password"))
		return
	}

	respondJSON(w, http.StatusOK, model.CreateResponse(u))
}
