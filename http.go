package phoval

import (
	"net/http"
	"phoval/mock"
	"phoval/service/notification"
	"time"
)

const (
	EnvProduction = "prod"
)

func NewHttpServer(env string, addr string, storage VerificationStorage, brand string) *Server {
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
		VerificationNotifier: getNotifier(env),
	}

	httpServer.Handler = server.Routes()

	return server
}

func getNotifier(env string) VerificationNotifier {
	if env == EnvProduction {
		return notification.AWSSESNotifier{}
	}

	return mock.smsNotifier{}
}
