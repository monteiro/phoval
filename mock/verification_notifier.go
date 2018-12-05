package mock

import (
	"monteiro/phoval/pkg/notification"
)

type VerificationNotification struct {
}

type MessageNotifier struct {
	Invoked bool
}

func (d MessageNotifier) Send(n notification.VerificationNotification) error {
	d.Invoked = true

	return nil
}
