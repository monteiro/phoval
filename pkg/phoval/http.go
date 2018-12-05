package phoval

import (
	"net/http"
	"time"
)

const (
	EnvProduction = "prod"
)

func NewHttpServer(addr string, storage VerificationStorage, brand string, notifier VerificationNotifier) *Server {
	httpServer := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 524288, //  limit the maximum header length to 0.5MB
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	server := &Server{
		Server:               httpServer,
		Storage:              storage,
		Brand:                brand,
		VerificationNotifier: notifier,
	}

	httpServer.Handler = server.Routes()

	return server
}
