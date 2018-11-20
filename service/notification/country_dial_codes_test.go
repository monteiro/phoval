package notification_test

import (
	"phoval/service/notification"
	"testing"
)

func TestValidateDialCode(t *testing.T) {

	d := notification.DialCode("+351")
	if !d.Valid() {
		t.Errorf("The dial code '%s' should be correct", (string(d)))
	}

	d2 := notification.DialCode("+999")
	if d2.Valid() {
		t.Errorf("The dial code '%s' should not exist", string(d2))
	}
}
