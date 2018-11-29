package phoval_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"phoval"
	"phoval/storage"
	"testing"
)

func TestHandleCreateVerification(t *testing.T) {
	method := "POST"
	url := "/phone/verification?country_code=351&phone_number=918888888"

	srv := phoval.NewHttpServer(":4000", storage.NewInMemoryStorage())
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Errorf("Error creating request")
	}

	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusCreated)
	}

	if w.Result().Header.Get("verification_id") == "" {
		t.Errorf("Response header should contain verification_id")
	}
}

func TestHandleVerification(t *testing.T) {
	id := "1"
	phoneNumber := "918888888"
	countryCode := "351"
	code := "123456"
	m := storage.NewInMemoryStorage()
	m.M[id] = phoval.PhoneCodeValidation{
		PhoneNumber: phoneNumber,
		CountryCode: countryCode,
		Code:        code,
	}

	method := "PUT"
	url := fmt.Sprintf("/phone/verification?country_code=%s&phone_number=%s&code=%s", countryCode, phoneNumber, code)

	srv := phoval.NewHttpServer(":4000", m)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Errorf("Error creating request")
	}

	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusNoContent)
	}
}
