package phoval

import (
	"fmt"
	"log"
	"monteiro/phoval/pkg/notification"
	"net/http"
	"strconv"
)

const (
	phoneNumberMandatoryErrorMessage = "'phone_number' is mandatory when creating a new verification"
	countryCodeMandatoryErrorMessage = "'country_code' is mandatory when creating a new verification"
	countryCodeNotValidErrorMessage  = "'country_code' is not valid"
	codeMandatoryErrorMessage        = "'code' is mandatory when creating a new verification"
)

func (s *Server) HandleCreateVerification() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locale, err := getQueryParam(r, "locale")
		if err != nil {
			locale = "en"
		}

		phoneNumber, err := getQueryParam(r, "phone_number")
		if err != nil {
			s.BadRequest(w, phoneNumberMandatoryErrorMessage)
			return
		}

		countryCode, err := getQueryParam(r, "country_code")
		if err != nil {
			s.BadRequest(w, countryCodeMandatoryErrorMessage)
			return
		}

		if !notification.DialCode(countryCode).Valid() {
			s.BadRequest(w, countryCodeNotValidErrorMessage)
			return
		}

		command := CreateVerificationCommand{
			CountryCode: countryCode,
			PhoneNumber: phoneNumber,
			Locale:      locale,
			From:        s.Brand,
		}

		resp, err := createVerificationCommandHandler(s.Storage, s.VerificationNotifier, s.NotificationRenderer, command)
		if err != nil {
			log.Fatalf("An internal server error has occurred: '%v'", err)
			s.InternalServerError(w)
			return
		}

		w.Header().Add("verification_id", resp.id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) HandleVerification() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phoneNumber, err := getQueryParam(r, "phone_number")
		if err != nil {
			s.BadRequest(w, phoneNumberMandatoryErrorMessage)
			return
		}

		countryCode, err := getQueryParam(r, "country_code")
		if err != nil {
			s.BadRequest(w, countryCodeMandatoryErrorMessage)
			return
		}

		if !notification.DialCode(countryCode).Valid() {
			s.BadRequest(w, countryCodeNotValidErrorMessage)
			return
		}

		code, err := getQueryParam(r, "code")
		if err != nil {
			s.BadRequest(w, codeMandatoryErrorMessage)
			return
		}

		_, err = strconv.Atoi(code)
		if err != nil {
			s.BadRequest(w, err.Error())
			return
		}

		command := ValidateCodeCommand{
			PhoneNumber: phoneNumber,
			CountryCode: countryCode,
			Code:        code,
		}

		err = VerifyCodeCommandHandler(s.Storage, command)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}

func getQueryParam(r *http.Request, index string) (string, error) {
	keys, ok := r.URL.Query()[index]
	if !ok || len(keys[0]) < 1 {
		return "", fmt.Errorf("param '%s' is missing", index)
	}

	return keys[0], nil
}
