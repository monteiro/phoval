package storage

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

type VerificationRepository interface {
	Validate(countryCode string, phoneNumber string, code int) error
	NewVerification(countryCode string, phoneNumber string) (string, error)
}
