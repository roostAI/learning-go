package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func TestInsertOneTask_841c76e883(t *testing.T) {
	// TODO: Initialize collection with your MongoDB collection

	// Test case: Inserting a valid task
	task1 := models.ToDoList{
		ID:     primitive.NewObjectID(),
		Task:   "Test Task 1",
		Status: "Pending",
	}
	insertOneTask(task1)
	insertedTask1, _ := collection.Find(context.Background(), bson.M{"task": "Test Task 1"})
	assert.NotNil(t, insertedTask1, "Failed to insert the task: ", task1)

	// Test case: Inserting a task with empty task field
	task2 := models.ToDoList{
		ID:     primitive.NewObjectID(),
		Task:   "",
		Status: "Pending",
	}
	insertOneTask(task2)
	insertedTask2, _ := collection.Find(context.Background(), bson.M{"task": ""})
	assert.Nil(t, insertedTask2, "Should not insert a task with an empty task field: ", task2)
}

func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
