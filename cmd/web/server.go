package main

import (
	"log"
	"net/http"
	"time"
)

func (app *App) RunServer() {
	srv := &http.Server{
		Addr:           app.Addr,
		MaxHeaderBytes: 524288, //  limit the maximum header length to 0.5MB
		Handler:        app.Routes(),
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	log.Printf("Starting server on %s", app.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
