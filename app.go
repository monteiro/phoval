package phoval

import "phoval/service/notification"

// VerifyNotification send a message to the user with the code
type PhoneVerification struct {
	// country code of the phone number
	CountryCode string
	// phone number to send the message
	PhoneNumber string
	// random code that is saved in the database
	Code string
}

type PhoneCodeValidation struct {
	// country code of the phone number
	CountryCode string
	// number without country code
	PhoneNumber string
	// code to validate the verification
	Code string
}

// Notifier that sends a specific message with the code to the user
type VerificationNotifier interface {
	Send(notification notification.VerificationNotification) error
}

type VerificationStorage interface {
	CreateVerification(v *PhoneVerification) (string, error)
	ValidateVerification(v *PhoneCodeValidation) error
}

type CreateVerificationCommand struct {
	// country code of the phone number
	CountryCode string
	// phone number to send the message
	PhoneNumber string
	// user's locale that defines the message language
	Locale string
	// recipient of the message
	From string
}

type CreateVerificationResponse struct {
	// verification id
	id string
}

type ValidateCodeCommand struct {
	// country code of the phone number
	CountryCode string
	// phone number to send the message
	PhoneNumber string
	// code to verify
	Code string
}
