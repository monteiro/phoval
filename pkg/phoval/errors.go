package phoval

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (s *Server) BadRequest(w http.ResponseWriter, desc string) {
	s.error(w, desc, http.StatusBadRequest)
}

func (s *Server) InternalServerError(w http.ResponseWriter) {
	s.error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (s *Server) NotFound(w http.ResponseWriter, desc string) {
	s.error(w, desc, http.StatusNotFound)
}

func (s *Server) error(w http.ResponseWriter, desc string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Code:        status,
		Description: desc,
	}); err != nil {
		checkError(err)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
