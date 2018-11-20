package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()
	mux.Post("/phone/verification", http.HandlerFunc(app.CreateNewVerification))
	mux.Put("/phone/verification", http.HandlerFunc(app.ValidatePhone))
	return LogRequest(mux)
}
