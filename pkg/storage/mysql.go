package storage

import (
	"2fa-api/pkg/generator"
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Validate validates phone number with a specific code that can be send by SMS
// countryCode country code that prefixes the phone number
// phoneNumber phone number without the country code
// code digits to verify the phone number
func (d *Database) Validate(countryCode string, phoneNumber string, code int) error {
	tx, err := d.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE verification SET verified_at = UTC_TIMESTAMP() WHERE country_code = ? and phone_number = ? AND code = ? AND verified_at IS NULL")
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(countryCode, phoneNumber, code)
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("Verification with phoneNumber: %v countryCode: %v and code: %v was not found", phoneNumber, countryCode, code)
	}

	return nil
}

// NewVerification create new phone number verification
func (d *Database) NewVerification(countryCode string, phoneNumber string) (string, error) {
	tx, err := d.Begin()
	if err != nil {
		return "", err
	}

	stmt, err := tx.Prepare("INSERT INTO verification(id, country_code, phone_number, code, created_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	countryCodeAsInt, err := strconv.Atoi(countryCode)
	if err != nil {
		return "", err
	}

	uuid := uuid.NewV4()
	code, err := generator.GenerateRandomDigits()
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(uuid.String(), countryCodeAsInt, phoneNumber, code, time.Now().UTC())
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return uuid.String(), nil
}
