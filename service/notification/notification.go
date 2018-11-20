package notification

type VerifyNotification struct {
	PhoneNumber string
	Message     string
	From        string
}

type Notifier interface {
	Send() error
}
