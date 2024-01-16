package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestTaskComplete_34e5efff97(t *testing.T) {
	// Test case 1: Valid task ID
	t.Run("Valid task ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/task/123", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/task/{id}", TaskComplete)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expected := `123`
		assert.Equal(t, expected, rr.Body.String())
	})

	// Test case 2: Invalid task ID
	t.Run("Invalid task ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/task/abc", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/task/{id}", TaskComplete)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expected := `abc`
		assert.Equal(t, expected, rr.Body.String())
	})
}
