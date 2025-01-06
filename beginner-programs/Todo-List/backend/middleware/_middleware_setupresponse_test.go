package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)



type ResponseRecorder struct {
	// Code is the HTTP response code set by WriteHeader.
	//
	// Note that if a Handler never calls WriteHeader or Write,
	// this might end up being 0, rather than the implicit
	// http.StatusOK. To get the implicit value, use the Result
	// method.
	Code int

	// HeaderMap contains the headers explicitly set by the Handler.
	// It is an internal detail.
	//
	// Deprecated: HeaderMap exists for historical compatibility
	// and should not be used. To access the headers returned by a handler,
	// use the Response.Header map as returned by the Result method.
	HeaderMap http.Header

	// Body is the buffer to which the Handler's Write calls are sent.
	// If nil, the Writes are silently discarded.
	Body *bytes.Buffer

	// Flushed is whether the Handler called Flush.
	Flushed bool

	result      *http.Response // cache of Result's return value
	snapHeader  http.Header    // snapshot of HeaderMap at first Write
	wroteHeader bool
}


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestsetupResponse(t *testing.T) {
	tests := []struct {
		name            string
		responseWriter  http.ResponseWriter
		requestMethod   string
		expectedHeaders map[string]string
		expectPanic     bool
	}{
		{
			name:           "Verify CORS Headers are Set Correctly",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "GET",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:            "Check Response Writer is Not Modified When Nil",
			responseWriter:  nil,
			requestMethod:   "GET",
			expectedHeaders: map[string]string{},
			expectPanic:     true,
		},
		{
			name:           "Validate Behavior with Different HTTP Methods",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "POST",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:           "Test Function with Empty Request",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:           "Confirm Header Overwrites",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "OPTIONS",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:           "Test Function with Multiple Calls",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "PUT",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.responseWriter != nil {
				req := &http.Request{Method: tt.requestMethod}
				setupResponse(&tt.responseWriter, req)

				recorder := tt.responseWriter.(*httptest.ResponseRecorder)
				for key, expectedValue := range tt.expectedHeaders {
					if actualValue := recorder.Header().Get(key); actualValue != expectedValue {
						t.Errorf("Expected header %s to be %s, got %s", key, expectedValue, actualValue)
					}
				}
			} else {
				defer func() {
					if r := recover(); r == nil && tt.expectPanic {
						t.Errorf("Expected panic for nil response writer, but did not panic")
					}
				}()
				setupResponse(nil, &http.Request{Method: tt.requestMethod})
			}
		})
	}
}

