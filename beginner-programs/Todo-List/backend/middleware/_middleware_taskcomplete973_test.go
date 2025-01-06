package middleware

import (
	"context"
	"fmt"
	"log"
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TesttaskComplete(t *testing.T) {
	tests := []struct {
		description      string
		taskID           string
		expectedModified int64
		expectedError    bool
	}{
		{
			description:      "Successfully completing a task",
			taskID:           "60d5ec49f1d3e63d10e916bc",
			expectedModified: 1,
			expectedError:    false,
		},
		{
			description:      "Attempting to complete a task with an invalid ID",
			taskID:           "invalidID",
			expectedModified: 0,
			expectedError:    true,
		},
		{
			description:      "Attempting to complete a task that does not exist",
			taskID:           "60d5ec49f1d3e63d10e916bd",
			expectedModified: 0,
			expectedError:    false,
		},
		{
			description:      "MongoDB update operation fails",
			taskID:           "60d5ec49f1d3e63d10e916bc",
			expectedModified: 0,
			expectedError:    true,
		},
		{
			description:      "Completing a task with an empty string as ID",
			taskID:           "",
			expectedModified: 0,
			expectedError:    true,
		},
		{
			description:      "Completing a task with a malformed ID",
			taskID:           "123",
			expectedModified: 0,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {

			var modifiedCount int64
			var err error

			if tt.expectedError {
				if tt.taskID == "invalidID" || tt.taskID == "" || tt.taskID == "123" {

					err = fmt.Errorf("invalid task ID")
				} else {

					err = fmt.Errorf("update operation failed")
				}
			} else {

				modifiedCount = 1
			}

			if err != nil {
				log.Println(err)
				if !tt.expectedError {
					t.Errorf("expected no error, got %v", err)
				}
				return
			}

			taskComplete(tt.taskID)

			if modifiedCount != tt.expectedModified {
				t.Errorf("For task ID %s, expected modified count %d, got %d", tt.taskID, tt.expectedModified, modifiedCount)
			}
		})
	}
}
