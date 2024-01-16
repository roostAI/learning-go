package middleware

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection 

func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

func TestDeleteOneTask(t *testing.T) {
	t.Run("Test with valid task ID", func(t *testing.T) {
		taskID := "60f7ea20b3b3d2c0018e47e7"
		deleteOneTask(taskID)

		id, _ := primitive.ObjectIDFromHex(taskID)
		filter := bson.M{"_id": id}
		err := collection.FindOne(context.Background(), filter).Err()

		assert.Equal(t, mongo.ErrNoDocuments, err, "Task was not deleted. Error:", err)
	})

	t.Run("Test with invalid task ID", func(t *testing.T) {
		taskID := "invalidTaskID"
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			}
		}()
		deleteOneTask(taskID)
	})
}
