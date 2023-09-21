package middleware

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func TestGetAllTask_dfbd966f8b(t *testing.T) {
	// Setup test database and collection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Error(err)
	}
	collection = client.Database("testDB").Collection("testCollection")

	// Test case 1: Check if the function returns results correctly
	results := getAllTask()
	if len(results) == 0 {
		t.Error("Expected to get at least one result, got none")
		t.Log("getAllTask() = ", results)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Check if the function handles errors correctly
	// For this, we will use a non-existent collection
	collection = client.Database("testDB").Collection("nonExistentCollection")
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected the function to panic, but it didn't")
		} else {
			t.Log("Test case 2 passed")
		}
	}()
	getAllTask()
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			panic(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	cur.Close(context.Background())
	return results
}
