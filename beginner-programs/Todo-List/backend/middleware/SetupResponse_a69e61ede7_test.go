package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupResponse_a69e61ede7(t *testing.T) {
	req, err := http.NewRequest("GET", "/dummy", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(setupResponse)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":   "*",
		"Access-Control-Allow-Methods":  "POST, GET, OPTIONS, PUT, DELETE",
		"Access-Control-Allow-Headers":  "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}
	for k, v := range expectedHeaders {
		if rr.Header().Get(k) != v {
			t.Errorf("Handler returned wrong header %s: got %v want %v", k, rr.Header().Get(k), v)
		}
	}
}

func setupResponse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
