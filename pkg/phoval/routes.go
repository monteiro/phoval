package phoval

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (s *Server) Routes() http.Handler {
	mux := pat.New()
	mux.Post("/phone/verification", http.HandlerFunc(LogRequest(s.HandleCreateVerification())))
	mux.Put("/phone/verification", http.HandlerFunc(LogRequest(s.HandleVerification())))

	return mux
}
