package notification

type VerificationNotification struct {
	CountryCode string
	// phone number to send the message
	PhoneNumber string
	// Message to send to the user
	Message string
	// Word that identifies who sends the message
	From string
}
