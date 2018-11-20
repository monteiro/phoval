package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (app *App) BadRequest(w http.ResponseWriter, desc string) {
	app.error(w, desc, http.StatusBadRequest)
}

func (app *App) NotFound(w http.ResponseWriter, desc string) {
	app.error(w, desc, http.StatusNotFound)
}

func (app *App) error(w http.ResponseWriter, desc string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Code:        status,
		Description: desc,
	}); err != nil {
		checkError(err)
	}
}
