package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
)

func TestCreateTask_eb4d396a39(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task", strings.NewReader(`{"title":"Test Task", "status": "pending"}`))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"title":"Test Task", "status": "pending"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
		t.Log("Success case passed")
	})

	t.Run("failure case", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task", strings.NewReader(`{"title":"", "status": "pending"}`))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		expected := `{"error":"Title is required"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
		t.Log("Failure case passed")
	})
}
