package notification

const (
	// Promotional are non-critical messages, such as marketing messages.
	// Amazon SNS optimizes the message delivery to incur the lowest cost.
	Promotional = "Promotional"

	// Transactional messages are critical messages that support
	// customer transactions, such as one-time passcodes for multi-factor authentication.
	// Amazon SNS optimizes the message delivery to achieve the highest reliability.
	Transactional = "Transactional"
)

// VerifyNotification send a message to the user with the code
type VerifyNotification struct {
	// phone number to send the message
	PhoneNumber string
	// Message to send to the user
	Message string
	// Word that identifies who sends the message
	From string
	// Promotional or Transactional
	Type string
}

// Notifier that sends a specific message with the code to the user
type Notifier interface {
	Send() error
}
