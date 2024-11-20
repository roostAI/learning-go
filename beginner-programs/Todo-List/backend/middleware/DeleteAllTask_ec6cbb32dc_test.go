package middleware

import (
	"context"
	"log"
	"testing"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mocked DeleteResult
type mockedDeleteResult struct {
	DeletedCount int64
}

func (m *mockedDeleteResult) DeletedCount() int64 {
	return m.DeletedCount
}

// Mocked Collection
type mockedCollection struct {
	DeleteManyFunc func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

func (m *mockedCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.DeleteManyFunc(ctx, filter, opts...)
}

func TestDeleteAllTask_ec6cbb32dc(t *testing.T) {
	// Mock the DeleteMany method
	collection := &mockedCollection{
		DeleteManyFunc: func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
			return &mongo.DeleteResult{DeletedCount: 5}, nil
		},
	}

	deletedCount := deleteAllTask(collection)

	if deletedCount != 5 {
		t.Errorf("Expected 5, but got %d", deletedCount)
	} else {
		t.Logf("Success: Expected 5 and got %d", deletedCount)
	}

	// Test the failure scenario
	collection = &mockedCollection{
		DeleteManyFunc: func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
			return nil, fmt.Errorf("Some error occurred")
		},
	}

	deletedCount = deleteAllTask(collection)

	if deletedCount != 0 {
		t.Errorf("Expected 0, but got %d", deletedCount)
	} else {
		t.Logf("Success: Expected 0 and got %d", deletedCount)
	}
}

func deleteAllTask(collection *mockedCollection) int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
