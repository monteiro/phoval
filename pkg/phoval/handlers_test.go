package phoval_test

import (
	"fmt"
	"monteiro/phoval/mock"
	"monteiro/phoval/pkg/phoval"
	"net/http"
	"net/http/httptest"
	"testing"

)

const anyPort = ":4000"

func TestHandleCreateVerification(t *testing.T) {
	url := "/phone/verification?country_code=351&phone_number=918888888"
	w := runTestHttpServer(t, mock.InMemoryVerificationStorage(), "POST", url)

	if w.Code != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: goot %v want %v", w.Code, http.StatusCreated)
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
	m := mock.InMemoryVerificationStorage()
	m.M[id] = phoval.PhoneCodeValidation{
		PhoneNumber: phoneNumber,
		CountryCode: countryCode,
		Code:        code,
	}

	url := fmt.Sprintf("/phone/verification?country_code=%s&phone_number=%s&code=%s", countryCode, phoneNumber, code)
	w := runTestHttpServer(t, m, "PUT", url)

	if w.Code != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusNoContent)
	}
}

var handleCreateVerificationArgsValidationTests = []struct {
	phoneNumber string
	countryCode string
}{
	{"", ""},
	{"91991999", ""},
	{"91991999", "3333333333"},
}

func TestHandleCreateVerificationValidationErrors(t *testing.T) {
	for _, vt := range handleCreateVerificationArgsValidationTests {
		t.Run(vt.phoneNumber+":"+vt.countryCode, func(t *testing.T) {
			url := fmt.Sprintf("/phone/verification?country_code=%s&phone_number=%s", vt.countryCode, vt.phoneNumber)
			w := runTestHttpServer(t, mock.InMemoryVerificationStorage(), "POST", url)
			if w.Code != http.StatusBadRequest {
				t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
			}
		})
	}
}

var handleVerificationArgsValidationTests = []struct {
	phoneNumber string
	countryCode string
	code        string
}{
	{"", "", ""},
	{"", "351", ""},
	{"91991999", "", "124556"},
	{"91991999", "351", ""},
}

func TestHandleVerificationValidationErrors(t *testing.T) {
	for _, vt := range handleVerificationArgsValidationTests {
		t.Run(vt.phoneNumber+":"+vt.countryCode+":"+vt.code, func(t *testing.T) {
			url := fmt.Sprintf("/phone/verification?country_code=%s&phone_number=%s&code=%s", vt.countryCode, vt.phoneNumber, vt.code)
			w := runTestHttpServer(t, mock.InMemoryVerificationStorage(), "PUT", url)
			if w.Code != http.StatusBadRequest {
				t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
			}
		})
	}
}

func runTestHttpServer(t *testing.T, v phoval.VerificationStorage, method string, url string) *httptest.ResponseRecorder {
	srv := phoval.NewHttpServer(anyPort, v, anyPort, mock.MessageNotifier{})
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Errorf("Error creating request")
	}

	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, req)

	return w
}
