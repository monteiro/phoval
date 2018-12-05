package phoval

import (
	"net/http"
)

type Server struct {
	*http.Server
	// address to listen for http requests
	Addr string
	// database
	Storage VerificationStorage
	// to be used in the message to send with the code in the message recipient
	Brand string
	// used to send a message
	VerificationNotifier VerificationNotifier
	// api key authorization
	ApiKey string
}
