package phoval

// VerifyNotification send a message to the user with the code
type PhoneVerification struct {
	// country code of the phone number
	CountryCode string
	// phone number to send the message
	PhoneNumber string
}

type PhoneCodeValidation struct {
	// country code of the phone number
	CountryCode string
	// number without country code
	PhoneNumber string
	// code to validate the verification
	Code string
}

type CreateVerificationCommand struct {
	// country code of the phone number
	CountryCode string
	// phone number to send the message
	PhoneNumber string
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

// Notifier that sends a specific message with the code to the user
type VerificationNotifier interface {
	Send(countryCode string, phoneNumber string) error
}

type VerificationStorage interface {
	CreateVerification(v *PhoneVerification) (string, error)
	ValidateVerification(v *PhoneCodeValidation) error
}
