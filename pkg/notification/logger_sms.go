package notification

import "log"

type LoggerSmsNotifier struct {
}

func (d LoggerSmsNotifier) Send(n VerificationNotification) error {
	log.Printf("SMS was sent: '%v'\n", n)
	return nil
}
