package middleware

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		return ""
	}

	return os.Getenv(key)
}

func TestInit_a8a462ccb8(t *testing.T) {
	var connectionString string

	// Get environment variable for connection string
	host := goDotEnvVariable("HOST")
	if "" == host {
		connectionString = "mongodb://localhost:27017"
	} else {
		connectionString = "mongodb://" + host + ":27017"
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		t.Log("Failed to connect to MongoDB: ", err)
		t.Fail()
	} else {
		t.Log("Successfully connected to MongoDB")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		t.Log("Failed to ping MongoDB: ", err)
		t.Fail()
	} else {
		t.Log("Successfully pinged MongoDB")
	}

	dbName := goDotEnvVariable("DB_NAME")
	collName := goDotEnvVariable("COLLECTION_NAME")
	collection := client.Database(dbName).Collection(collName)

	if collection == nil {
		t.Log("Failed to create collection instance")
		t.Fail()
	} else {
		t.Log("Successfully created collection instance")
	}
}
