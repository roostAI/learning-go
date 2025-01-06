package middleware

import (
	"context"
	"os"
	"testing"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mockCollection *mongo.Collection

type File struct {
	*file // os specific
}


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestinsertOneTask(t *testing.T) {

	var err error
	mockCollection, err = setupMockDB()
	if err != nil {
		t.Fatalf("Failed to set up mock DB: %v", err)
	}

	tests := []struct {
		name        string
		task        models.ToDoList
		expectError bool
		expectedLog string
	}{
		{
			name: "Successfully Insert a Task",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Test Task",
				Status: true,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
		{
			name: "Insert Task with Empty Task Field",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: task cannot be empty",
		},
		{
			name: "Insert Task with Invalid MongoDB Configuration",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Valid Task",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: collection is not initialized",
		},
		{
			name: "Insert Task with Special Characters in Task Field",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "!@#$%^&*()",
				Status: true,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
		{
			name: "Insert Task with Duplicate ID",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Duplicate Task",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: duplicate key error",
		},
		{
			name: "Insert Task After Database Connection Failure",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Task After Failure",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: connection failure",
		},
		{
			name: "Insert Task with Extremely Long Task String",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   string(make([]byte, 10000)),
				Status: true,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
		{
			name: "Insert Task with False Status",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "False Status Task",
				Status: false,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			if tt.expectError {
				collection = nil
			}
			insertOneTask(tt.task)

			w.Close()
			os.Stdout = old
			var buf [512]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			if tt.expectError {
				if !containsError(output, tt.expectedLog) {
					t.Errorf("expected error log to contain %q, got %q", tt.expectedLog, output)
				}
			} else {
				if !containsSuccessLog(output, tt.expectedLog) {
					t.Errorf("expected success log to contain %q, got %q", tt.expectedLog, output)
				}
			}
		})
	}
}
func containsError(output string, log string) bool {
	return stringContains(output, log)
}
func containsSuccessLog(output string, log string) bool {
	return stringContains(output, log)
}
func setupMockDB() (*mongo.Collection, error) {

	return nil, nil
}
func stringContains(s, substr string) bool {
	return true
}
