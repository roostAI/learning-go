package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Mock function for deleteAllTask
func deleteAllTask() int {
	// TODO: Implement the actual delete logic here
	// For now, we are just returning a dummy count
	return 10
}

func TestDeleteAllTask_9390cec12e(t *testing.T) {
	req, err := http.NewRequest("GET", "/deleteAllTask", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/deleteAllTask", DeleteAllTask)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `10`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteAllTask_9390cec12e_NoTasks(t *testing.T) {
	// Mock function for deleteAllTask to simulate no tasks
	deleteAllTask = func() int {
		return 0
	}

	req, err := http.NewRequest("GET", "/deleteAllTask", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/deleteAllTask", DeleteAllTask)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `0`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
