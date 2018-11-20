package main

import (
	"fmt"
	"net/http"
	"phoval/service/notification"
	"strconv"
)

// CreateNewVerification Create a shortener
func (app *App) CreateNewVerification(w http.ResponseWriter, r *http.Request) {

	phoneNumber, err := getQueryParam(r, "phone_number")
	if err != nil {
		app.BadRequest(w, "'phone_number' is mandatory when creating a new verification")
		return
	}

	countryCode, err := getQueryParam(r, "country_code")
	if err != nil {
		app.BadRequest(w, "'country_code' is mandatory when creating a new verification")
		return
	}

	if !notification.DialCode(countryCode).Valid() {
		app.BadRequest(w, "'country_code' is not valid")
		return
	}

	// validate phone number and country code
	id, err := app.Database.NewVerification(countryCode, phoneNumber)

	w.Header().Add("verification_id", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// ValidatePhone validates phone number with a code
func (app *App) ValidatePhone(w http.ResponseWriter, r *http.Request) {
	phoneNumber, err := getQueryParam(r, "phone_number")
	if err != nil {
		app.BadRequest(w, "'phone_number' is mandatory when creating a new verification")
		return
	}

	countryCode, err := getQueryParam(r, "country_code")
	if err != nil {
		app.BadRequest(w, "'country_code' is mandatory when creating a new verification")
		return
	}

	if !notification.DialCode(countryCode).Valid() {
		app.BadRequest(w, "'country_code' is not valid")
		return
	}

	code, err := getQueryParam(r, "code")
	if err != nil {
		app.BadRequest(w, "'code' is mandatory when creating a new verification")
		return
	}

	c, err := strconv.Atoi(code)
	if err != nil {
		app.BadRequest(w, err.Error())
		return
	}
	// update table and generate the verified_at date
	err = app.Database.Validate(countryCode, phoneNumber, c)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func getQueryParam(r *http.Request, index string) (string, error) {
	keys, ok := r.URL.Query()[index]
	if !ok || len(keys[0]) < 1 {
		return "", fmt.Errorf("Param '%s' is missing", index)
	}

	return keys[0], nil
}
