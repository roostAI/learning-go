package router

import (
	"net/http"
	"testing"
	"github.com/gorilla/mux"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/middleware"
)




type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestRouter(t *testing.T) {
	tests := []struct {
		name            string
		method          string
		path            string
		expectedHandler http.HandlerFunc
	}{
		{
			name:            "Verify Router Initialization",
			method:          "GET",
			path:            "/todo",
			expectedHandler: nil,
		},
		{
			name:            "Check Route Handlers for GET /todo",
			method:          "GET",
			path:            "/todo",
			expectedHandler: middleware.GetAllTask,
		},
		{
			name:            "Check Route Handlers for POST /todo",
			method:          "POST",
			path:            "/todo",
			expectedHandler: middleware.CreateTask,
		},
		{
			name:            "Check Route Handlers for PUT /todo/{id}",
			method:          "PUT",
			path:            "/todo/{id}",
			expectedHandler: middleware.TaskComplete,
		},
		{
			name:            "Check Route Handlers for PUT /todo/undoTask/{id}",
			method:          "PUT",
			path:            "/todo/undoTask/{id}",
			expectedHandler: middleware.UndoTask,
		},
		{
			name:            "Check Route Handlers for DELETE /todo/{id}",
			method:          "DELETE",
			path:            "/todo/deleteTask/{id}",
			expectedHandler: middleware.DeleteTask,
		},
		{
			name:            "Check Route Handlers for DELETE /todo",
			method:          "DELETE",
			path:            "/todo",
			expectedHandler: middleware.DeleteAllTask,
		},
		{
			name:            "Validate OPTIONS Method for /todo",
			method:          "OPTIONS",
			path:            "/todo",
			expectedHandler: nil,
		},
		{
			name:            "Check for NotFoundHandler Configuration",
			method:          "GET",
			path:            "/undefined",
			expectedHandler: nil,
		},
		{
			name:            "Check for MethodNotAllowedHandler Configuration",
			method:          "POST",
			path:            "/todo",
			expectedHandler: nil,
		},
	}

	router := Router()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedHandler == nil {

				if router == nil {
					t.Error("Expected router to be initialized, got nil")
				}
				return
			}

			route := router.Get(tt.path)
			if route == nil {
				t.Errorf("No route found for path: %s", tt.path)
				return
			}

			if route.GetHandler() != tt.expectedHandler {
				t.Errorf("Expected handler for %s %s to be %v, got %v", tt.method, tt.path, tt.expectedHandler, route.GetHandler())
			}

			if tt.method == "OPTIONS" {
				if !route.Methods("OPTIONS") {
					t.Errorf("Expected OPTIONS method to be supported for %s", tt.path)
				}
			}

			if tt.name == "Check for NotFoundHandler Configuration" {
				if router.NotFoundHandler == nil {
					t.Error("Expected NotFoundHandler to be configured, got nil")
				}
			}

			if tt.name == "Check for MethodNotAllowedHandler Configuration" {
				if router.MethodNotAllowedHandler == nil {
					t.Error("Expected MethodNotAllowedHandler to be configured, got nil")
				}
			}
		})
	}
}
