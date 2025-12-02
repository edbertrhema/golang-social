package main

import (
	"net/http"
)

func (app *application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "UP",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err == nil {
		app.internalServerError(w, r, err)
	}

}
