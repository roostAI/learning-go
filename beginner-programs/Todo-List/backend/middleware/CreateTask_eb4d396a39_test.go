package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask_eb4d396a39(t *testing.T) {
	// Test Case 1: Check for successful task creation
	t.Run("success case", func(t *testing.T) {
		task := models.ToDoList{
			Task: "Test Task",
		}
		jsonTask, _ := json.Marshal(task)
		req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonTask))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var responseTask models.ToDoList
		_ = json.NewDecoder(rr.Body).Decode(&responseTask)
		assert.Equal(t, task.Task, responseTask.Task)
	})

	// Test Case 2: Check for unsuccessful task creation when body is empty
	t.Run("failure case", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
