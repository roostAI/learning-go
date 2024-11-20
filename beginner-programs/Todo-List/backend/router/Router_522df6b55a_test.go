package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todo", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/todo", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/todo/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/todo/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/todo/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/todo", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}

func TestRouter_522df6b55a(t *testing.T) {
	r := Router()

	// Test case 1: GET /todo
	req, _ := http.NewRequest("GET", "/todo", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test case 2: POST /todo
	req, _ = http.NewRequest("POST", "/todo", nil) // TODO: add a body to the request
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
