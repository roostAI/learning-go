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

	ts := httptest.NewServer(r)
	defer ts.Close()

	// Test GetAllTask
	resp, err := http.Get(ts.URL + "/todo")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code for /todo is wrong. Have: %d, want: %d.", resp.StatusCode, http.StatusOK)
	}

	// Test CreateTask
	resp, err = http.Post(ts.URL+"/todo", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code for /todo is wrong. Have: %d, want: %d.", resp.StatusCode, http.StatusOK)
	}

	// TODO: Add more tests for TaskComplete, UndoTask, DeleteTask, DeleteAllTask
}
