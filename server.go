package phoval

import (
	"net/http"
	"phoval/service/notification"
)

type Server struct {
	*http.Server
	// address to listen for http requests
	Addr string
	// database
	Storage VerificationStorage
	// Notifier to send SMS or Emails
	Notifier notification.Notifier
}
