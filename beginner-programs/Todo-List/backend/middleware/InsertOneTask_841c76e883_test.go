package middleware

import (
	"context"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/assert"
)

// TODO: Replace with your MongoDB connection string
var connectionString = "mongodb://localhost:27017"

// TODO: Replace with your database name
var dbName = "test_db"

// TODO: Replace with your collection name
var collectionName = "test_collection"

var collection *mongo.Collection

func init() {
	// Load .env file if exists
	godotenv.Load()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collectionName)
}

func insertOneTask(task models.ToDoList) error {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		return err
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	return nil
}

func TestInsertOneTask_841c76e883(t *testing.T) {
	// Test case 1: Insert a valid task
	task1 := models.ToDoList{
		Task:     "Test Task 1",
		Status:   "Pending",
		Comments: "This is a test task",
	}
	err := insertOneTask(task1)
	assert.Nil(t, err)
	t.Log("Success: Inserted a valid task")

	// Test case 2: Insert a task without task name
	task2 := models.ToDoList{
		Status:   "Pending",
		Comments: "This is a test task without task name",
	}
	err = insertOneTask(task2)
	assert.NotNil(t, err)
	if err == nil {
		t.Error("Failure: Inserted a task without task name")
	} else {
		t.Log("Success: Did not allow insertion of task without task name")
	}
}
