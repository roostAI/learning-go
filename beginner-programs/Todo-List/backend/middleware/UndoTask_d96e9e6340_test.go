package middleware

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func undoTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("modified count: ", result.ModifiedCount)
}

func TestUndoTask_d96e9e6340(t *testing.T) {
	// TODO: Initialize the MongoDB connection and set the collection
	// collection = ...

	// Test case 1: Valid task id
	taskID := "60b77d2f4f88d445e0f8e3e6" // TODO: Replace with a valid task id from your MongoDB
	undoTask(taskID)
	// Assuming the task with the given id exists and its status is true, the modified count should be 1
	// TODO: Check the status of the task in the database, it should be false now

	// Test case 2: Invalid task id
	taskID = "invalid-id"
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	undoTask(taskID)
	// The code should panic because the task id is not a valid hex
}
