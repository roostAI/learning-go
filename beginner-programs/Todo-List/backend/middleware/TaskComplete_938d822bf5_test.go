package middleware

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func TestTaskComplete_938d822bf5(t *testing.T) {
	// TODO: Replace with your MongoDB URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Error("Failed to connect to MongoDB", err)
		return
	}
	collection = client.Database("test").Collection("tasks")

	// Test case 1: Valid task ID
	taskID := "60b1f9c9d82b8c13a3753d42" // TODO: Replace with a valid task ID in your database
	taskComplete(taskID)
	assert.Nil(t, err, "Test case 1: Passed")

	// Test case 2: Invalid task ID
	taskID = "invalid"
	defer func() {
		if r := recover(); r != nil {
			t.Log("Test case 2: Passed")
		} else {
			t.Error("Test case 2: Failed. Expected panic for invalid task ID, but function did not panic.")
		}
	}()
	taskComplete(taskID)
}
