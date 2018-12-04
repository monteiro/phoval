package mock

import (
	"errors"
	"phoval"

	"github.com/satori/go.uuid"
)

type InMemoryStorage struct {
	M map[string]phoval.PhoneCodeValidation
}

func (s *InMemoryStorage) CreateVerification(v *phoval.PhoneVerification) (string, error) {
	id := uuid.NewV4().String()
	s.M[id] = phoval.PhoneCodeValidation{
		CountryCode: v.CountryCode,
		PhoneNumber: v.PhoneNumber,
		Code:        v.Code,
	}

	return id, nil
}

func (s *InMemoryStorage) ValidateVerification(v *phoval.PhoneCodeValidation) error {

	for _, c := range s.M {
		if c.CountryCode == v.CountryCode && c.PhoneNumber == v.PhoneNumber {
			return nil
		}
	}

	return errors.New("no validation request found")
}

func InMemoryVerificationStorage() *InMemoryStorage {
	return &InMemoryStorage{
		M: make(map[string]phoval.PhoneCodeValidation),
	}
}
