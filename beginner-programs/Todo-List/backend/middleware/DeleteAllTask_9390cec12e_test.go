package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAllTask_9390cec12e(t *testing.T) {
	// TODO: Setup the database connection and add some tasks before running the test

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/deleteAllTasks", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteAllTask)

	// Create a router for middleware because DeleteAllTask requires mux.Vars to get params
	r := mux.NewRouter()
	r.Handle("/deleteAllTasks", handler)

	// Run the handler and check if the status code is what we expect
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Check the response body
	expected := `0`
	assert.Equal(t, expected, rr.Body.String(), "handler returned unexpected body")
}
