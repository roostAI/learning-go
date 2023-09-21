package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTask_9ddf916cf4(t *testing.T) {
	// Test Case 1: Valid ID
	req, err := http.NewRequest("DELETE", "/task/validID", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/task/{id}", DeleteTask)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	expected := `{"id":"validID"}`
	assert.Equal(t, expected, rr.Body.String(), "handler returned unexpected body")
	t.Log("Test Case 1 Passed")

	// Test Case 2: Invalid ID
	req, err = http.NewRequest("DELETE", "/task/invalidID", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router = mux.NewRouter()
	router.HandleFunc("/task/{id}", DeleteTask)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "handler returned wrong status code")

	expected = `{"error":"Task not found"}`
	assert.Equal(t, expected, rr.Body.String(), "handler returned unexpected body")
	t.Log("Test Case 2 Passed")
}
