package phoval

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pattern := `%s - "%s %s %s"`
		log.Printf(pattern, r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		h(w, r)
	}
}

func ApiKeyAuthorization(h http.HandlerFunc, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerApiKey := r.Header.Get("Authorization")
		if headerApiKey != apiKey {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			status := http.StatusUnauthorized
			w.WriteHeader(status)

			if err := json.NewEncoder(w).Encode(ErrorResponse{
				Code:        status,
				Description: fmt.Sprintf("api key '%s' not authorized", headerApiKey),
			}); err != nil {
				checkError(err)
			}

			return
		}

		h(w, r)
	}
}
