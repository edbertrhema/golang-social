package main

import (
	"database/sql"
	"errors"
	"net/http"
	"social/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreateUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err := Validate.Struct(&payload)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	user := store.User{
		Username: payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := app.store.Users.Create(ctx, &user); err != nil {
		switch {
		case errors.Is(err, store.ErrDuplicateKey):
			app.duplicateError(w, r, err)
		default:
			app.internalServerError(w, r, err)

		}
		return
	}

	app.jsonResponse(w, http.StatusCreated, user)
}

func (app *application) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	result, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
