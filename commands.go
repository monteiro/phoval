package phoval

func createVerificationCommandHandler(s VerificationStorage, c CreateVerificationCommand) (CreateVerificationResponse, error) {
	id, err := s.CreateVerification(&PhoneVerification{
		CountryCode: c.CountryCode,
		PhoneNumber: c.PhoneNumber,
	})

	if err != nil {
		return CreateVerificationResponse{}, err
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
		return err
	}

	return nil
}
