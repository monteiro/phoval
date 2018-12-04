package mysql

import (
	"database/sql"
	"fmt"
	"phoval"
	"phoval/pkg/generator"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

const insertVerification = `INSERT INTO verification(id, country_code, phone_number, code, created_at) VALUES(?, ?, ?, ?, ?)`
const validateVerification = `UPDATE verification SET verified_at = UTC_TIMESTAMP() WHERE country_code = ? and phone_number = ? AND code = ? AND verified_at IS NULL`

type VerificationStorage struct {
	DB *sql.DB
}

func (s *VerificationStorage) CreateVerification(v *phoval.PhoneVerification) (string, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return "", err
	}

	stmt, err := tx.Prepare(insertVerification)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	countryCode, err := strconv.Atoi(v.CountryCode)
	if err != nil {
		return "", err
	}

	id := uuid.NewV4()
	code, err := generator.GenerateRandomDigits()
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(id.String(), countryCode, v.PhoneNumber, code, time.Now().UTC())
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return id.String(), nil
}

func (s *VerificationStorage) ValidateVerification(v *phoval.PhoneCodeValidation) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(validateVerification)
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(v.CountryCode, v.PhoneNumber, v.Code)
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf(
			"verification with phoneNumber: %v countryCode: %v and code: %v was not found", v.PhoneNumber, v.CountryCode, v.Code)
	}

	return nil
}
