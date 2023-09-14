package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAllTask_09b678baa9(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetAllTask)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		t.Log("Method Arguments: ", req)
	} else {
		t.Log("TestGetAllTask_09b678baa9 passed")
	}
}

func TestGetAllTask_09b678baa9_EmptyResult(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Empty the task list before running this test

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetAllTask)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		t.Log("Method Arguments: ", req)
	} else if rr.Body.String() != "[]" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "[]")
		t.Log("Method Arguments: ", req)
	} else {
		t.Log("TestGetAllTask_09b678baa9_EmptyResult passed")
	}
}
