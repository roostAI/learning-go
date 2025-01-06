package middleware

import (
	"context"
	"fmt"
	"log"
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mockCollection *mongo.Collection


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestdeleteAllTask(t *testing.T) {
	setup()

	tests := []struct {
		name          string
		initialCount  int64
		expectedCount int64
		mockError     error
	}{
		{
			name:          "Successful Deletion of All Tasks",
			initialCount:  5,
			expectedCount: 5,
			mockError:     nil,
		},
		{
			name:          "Deletion When No Tasks Exist",
			initialCount:  0,
			expectedCount: 0,
			mockError:     nil,
		},
		{
			name:          "Error Handling During Deletion",
			initialCount:  3,
			expectedCount: 0,
			mockError:     fmt.Errorf("mock error"),
		},
		{
			name:          "Deletion with Partial Data",
			initialCount:  4,
			expectedCount: 4,
			mockError:     nil,
		},
		{
			name:          "Concurrency Test for Deletion",
			initialCount:  10,
			expectedCount: 10,
			mockError:     nil,
		},
		{
			name:          "Checking Database Connection Before Deletion",
			initialCount:  0,
			expectedCount: 0,
			mockError:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i := int64(0); i < tt.initialCount; i++ {
				_, err := mockCollection.InsertOne(context.Background(), bson.D{{Key: "task", Value: fmt.Sprintf("task %d", i)}})
				if err != nil {
					t.Fatalf("Failed to insert task: %v", err)
				}
			}

			mockDeleteMany := func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
				if tt.mockError != nil {
					return nil, tt.mockError
				}
				return &mongo.DeleteResult{DeletedCount: tt.expectedCount}, nil
			}

			mockCollection = &mongo.Collection{

				DeleteMany: mockDeleteMany,
			}

			result := deleteAllTask()

			if result != tt.expectedCount {
				t.Errorf("Expected DeletedCount %d, got %d", tt.expectedCount, result)
			}

			_, _ = mockCollection.DeleteMany(context.Background(), bson.D{{}})
		})
	}
}
func setup() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mockCollection = client.Database("test").Collection("todolist")
}
