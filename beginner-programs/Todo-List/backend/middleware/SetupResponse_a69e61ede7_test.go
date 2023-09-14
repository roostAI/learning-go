package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func TestSetupResponse_a69e61ede7(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
	})

	handler.ServeHTTP(rr, req)

	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":   "*",
		"Access-Control-Allow-Methods":  "POST, GET, OPTIONS, PUT, DELETE",
		"Access-Control-Allow-Headers":  "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}

	for key, value := range expectedHeaders {
		if rr.Header().Get(key) != value {
			t.Errorf("header %s = %s, want %s", key, rr.Header().Get(key), value)
		}
	}
}

func TestSetupResponse_a69e61ede7_Failure(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
	})

	handler.ServeHTTP(rr, req)

	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":   "WrongValue",
		"Access-Control-Allow-Methods":  "WrongValue",
		"Access-Control-Allow-Headers":  "WrongValue",
	}

	for key, value := range expectedHeaders {
		if rr.Header().Get(key) == value {
			t.Errorf("header %s = %s, want not %s", key, rr.Header().Get(key), value)
		}
	}
}
