package phoval

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (s *Server) Routes() http.Handler {
	mux := pat.New()
	mux.Post("/phone/verification", http.HandlerFunc(LogRequest(ApiKeyAuthorization(s.HandleCreateVerification(), s.ApiKey))))
	mux.Put("/phone/verification", http.HandlerFunc(LogRequest(ApiKeyAuthorization(s.HandleVerification(), s.ApiKey))))

	return mux
}
