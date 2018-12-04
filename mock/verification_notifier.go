package mock

import (
	"log"
	"phoval/service/notification"
)

type VerificationNotification struct {
}

type SmsNotifier struct {
}

func (d SmsNotifier) Send(n notification.VerificationNotification) error {
	log.Printf("SMS was sent: '%v'\n", n)
	return nil
}
