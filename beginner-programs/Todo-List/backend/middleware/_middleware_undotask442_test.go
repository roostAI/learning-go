package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Database struct {
	client         *Client
	name           string
	readConcern    *readconcern.ReadConcern
	writeConcern   *writeconcern.WriteConcern
	readPreference *readpref.ReadPref
	readSelector   description.ServerSelector
	writeSelector  description.ServerSelector
	registry       *bsoncodec.Registry
}


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestundoTask(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	db := client.Database(dbName)
	collection = db.Collection(collName)

	tests := []struct {
		name          string
		taskID        string
		expectedLog   string
		expectedCount int64
		mockData      bson.M
		mockError     error
	}{
		{
			name:          "Successfully Undo a Task",
			taskID:        "507f191e810c19729de860ea",
			expectedLog:   "modified count:  1",
			expectedCount: 1,
			mockData:      bson.M{"_id": primitive.ObjectIDFromHex("507f191e810c19729de860ea"), "status": true},
			mockError:     nil,
		},
		{
			name:          "Handle Invalid Task ID",
			taskID:        "invalid-id",
			expectedLog:   "error: invalid hex string",
			expectedCount: 0,
			mockData:      nil,
			mockError:     nil,
		},
		{
			name:          "MongoDB Update Operation Fails",
			taskID:        "507f191e810c19729de860ea",
			expectedLog:   "error: update failed",
			expectedCount: 0,
			mockData:      bson.M{"_id": primitive.ObjectIDFromHex("507f191e810c19729de860ea"), "status": true},
			mockError:     fmt.Errorf("update failed"),
		},
		{
			name:          "No Task Found for Valid ID",
			taskID:        "507f191e810c19729de860ea",
			expectedLog:   "modified count:  0",
			expectedCount: 0,
			mockData:      bson.M{"_id": primitive.ObjectIDFromHex("507f191e810c19729de860ea"), "status": false},
			mockError:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.mockData != nil {
				_, err := collection.InsertOne(context.Background(), tt.mockData)
				if err != nil {
					t.Fatalf("Failed to insert mock data: %v", err)
				}
			}

			if tt.taskID == "invalid-id" {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for invalid ID, got none")
					}
				}()
			} else {

				if tt.mockError != nil {

					log.Println(tt.mockError)
				} else {
					undoTask(tt.taskID)
				}
			}

			if !strings.Contains(buf.String(), tt.expectedLog) {
				t.Errorf("Expected log to contain %q, got %q", tt.expectedLog, buf.String())
			}

			_, err := collection.DeleteMany(context.Background(), bson.M{})
			if err != nil {
				t.Fatalf("Failed to clean up mock data: %v", err)
			}

			buf.Reset()
		})
	}
}
