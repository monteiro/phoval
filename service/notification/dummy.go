package notification

import (
	"log"
)

type DummyNotification struct {
}

func (d DummyNotification) Send(n VerificationNotification) error {
	log.Printf("SMS was sent: '%v'\n", n)
	return nil
}
