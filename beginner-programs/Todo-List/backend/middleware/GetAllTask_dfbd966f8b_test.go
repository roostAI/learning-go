package middleware

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/stretchr/testify/assert"
)

var collection *mongo.Collection

func TestGetAllTask_dfbd966f8b(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		// TODO: Replace with your own valid MongoDB collection
		client, _ := mongo.NewClient(options.Client().ApplyURI("<MongoDB_URI>"))
		collection = client.Database("<database_name>").Collection("<collection_name>")
		results := getAllTask()
		if len(results) == 0 {
			t.Error("No data found in collection")
			t.Log("getAllTask() = ", results)
		} else {
			t.Log("Success! Data found in collection")
		}
	})

	t.Run("Failure case", func(t *testing.T) {
		// TODO: Replace with your own invalid MongoDB collection
		client, _ := mongo.NewClient(options.Client().ApplyURI("<MongoDB_URI>"))
		collection = client.Database("<database_name>").Collection("<invalid_collection_name>")
		results := getAllTask()
		assert.Nil(t, results, "Expected nil, but got data")
		if len(results) != 0 {
			t.Error("Data found in invalid collection")
			t.Log("getAllTask() = ", results)
		} else {
			t.Log("Success! No data found in invalid collection")
		}
	})
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		t.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			t.Fatal(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		t.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}
