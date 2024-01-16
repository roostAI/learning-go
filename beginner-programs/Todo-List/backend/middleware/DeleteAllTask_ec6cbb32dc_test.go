package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func TestDeleteAllTask_ec6cbb32dc(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") 
	client, err := mongo.Connect(context.Background(), clientOptions) 
	assert.NoError(t, err, "Failed to connect to database")

	collection = client.Database("test").Collection("tasks")

	// Test case 1: Delete all tasks, expecting success
	deletedCount := deleteAllTask()
	assert.NotEqual(t, -1, deletedCount, "Failed to delete all tasks")

	// Test case 2: Delete all tasks when there are no tasks, expecting success
	deletedCount = deleteAllTask()
	assert.Equal(t, 0, deletedCount, "Failed to handle case when there are no tasks to delete")
}

func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		return -1
	}

	return d.DeletedCount
}
