package phoval

import (
	"log"
	"monteiro/phoval/pkg/generator"
	"monteiro/phoval/pkg/notification"
)

func createVerificationCommandHandler(s VerificationStorage, v VerificationNotifier, n NotificationRenderer, c CreateVerificationCommand) (CreateVerificationResponse, error) {
	r := CreateVerificationResponse{}

	code, err := generator.GenerateRandomDigits()
	if err != nil {
		return r, err
	}

	id, err := s.CreateVerification(&PhoneVerification{
		CountryCode: c.CountryCode,
		PhoneNumber: c.PhoneNumber,
		Code:        code,
	})
	if err != nil {
		return r, err
	}

	m, err := n.Render(c.Locale, code)
	if err != nil {
		return r, err
	}

	notif := notification.VerificationNotification{
		CountryCode: c.CountryCode,
		PhoneNumber: c.PhoneNumber,
		From:        c.From,
		Message:     m,
	}

	if err := v.Send(notif); err != nil {
		log.Fatalf("error sending message: '%v'", n)
		return r, err
	}

	return CreateVerificationResponse{
		id: id,
	}, nil
}

func VerifyCodeCommandHandler(s VerificationStorage, c ValidateCodeCommand) error {
	err := s.ValidateVerification(&PhoneCodeValidation{
		CountryCode: c.CountryCode,
		PhoneNumber: c.PhoneNumber,
		Code:        c.Code,
	})

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
