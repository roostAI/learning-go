package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func TestDeleteOneTask_22ade46ed5(t *testing.T) {
	// TODO: Change the connection string as per your configuration
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	assert.NoError(t, err)

	// TODO: Change the database and collection name as per your configuration
	collection = client.Database("test").Collection("tasks")

	// Test case 1: Valid task id
	taskID := "60b6c2f9f8a9d3f8f6e1727d" // TODO: Replace with a valid task id in your database
	deleteOneTask(t, taskID)
	_, err = collection.Find(context.Background(), bson.M{"_id": taskID})
	assert.Error(t, err, "Expected task to be deleted, but it still exists")

	// Test case 2: Invalid task id
	taskID = "invalid"
	deleteOneTask(t, taskID)
	_, err = collection.Find(context.Background(), bson.M{"_id": taskID})
	assert.Error(t, err, "Expected error due to invalid task id, but got none")
}

func deleteOneTask(t *testing.T, task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	assert.NoError(t, err)
	assert.NotZero(t, d.DeletedCount, "Expected at least one document to be deleted")
}
