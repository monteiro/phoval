package phoval_test

import (
	"fmt"
	"monteiro/phoval/messages"
	"monteiro/phoval/mock"
	"monteiro/phoval/pkg/phoval"
	"net/http"
	"net/http/httptest"
	"testing"
)

const anyPort = ":4000"
const apiKey = "testapikey"

func TestUnauthorized(t *testing.T) {
	url := "/phone/verification?country_code=351&phone_number=918888888"
	s := runTestHttpServer(mock.InMemoryVerificationStorage(), apiKey)
	w := makeRequest(t, "POST", url, "wrongapikey", s)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusUnauthorized)
	}
}

func TestHandleCreateVerification(t *testing.T) {
	url := "/phone/verification?country_code=351&phone_number=918888888"
	s := runTestHttpServer(mock.InMemoryVerificationStorage(), apiKey)
	w := makeRequest(t, "POST", url, apiKey, s)

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
	m := mock.InMemoryVerificationStorage()
	m.M[id] = phoval.PhoneCodeValidation{
		PhoneNumber: phoneNumber,
		CountryCode: countryCode,
		Code:        code,
	}

	url := fmt.Sprintf("/phone/verification?country_code=%s&phone_number=%s&code=%s", countryCode, phoneNumber, code)
	s := runTestHttpServer(m, apiKey)
	w := makeRequest(t, "PUT", url, apiKey, s)

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
			s := runTestHttpServer(mock.InMemoryVerificationStorage(), apiKey)
			w := makeRequest(t, "POST", url, apiKey, s)

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
			s := runTestHttpServer(mock.InMemoryVerificationStorage(), apiKey)
			w := makeRequest(t, "PUT", url, apiKey, s)

			if w.Code != http.StatusBadRequest {
				t.Errorf("Handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
			}
		})
	}
}

func runTestHttpServer(v phoval.VerificationStorage, apiKey string) *phoval.Server {
	return phoval.NewHttpServer(anyPort, v, anyPort, mock.MessageNotifier{}, apiKey, messages.TemplateFolderRender{
		TemplateFolder: "../../messages",
	})
}

func makeRequest(t *testing.T, method string, url string, apiKey string, srv *phoval.Server) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", apiKey)
	if err != nil {
		t.Errorf("error creating request")
	}
	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, req)
	return w
}
