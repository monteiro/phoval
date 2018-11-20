package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()
	mux.Post("/phone/verification/status", http.HandlerFunc(app.CreateNewVerification))
	mux.Put("/phone/verification/status", http.HandlerFunc(app.ValidatePhone))
	return LogRequest(mux)
}
